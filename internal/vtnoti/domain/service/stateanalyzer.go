package service

import (
	"context"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/blockchain/contracts/kyberswapbsc"
	"github.com/thanhpp/gopher/pkg/blockchain/onchain"
	"github.com/thanhpp/gopher/pkg/crawlers/bscscan"
)

type StateAnalyzer struct {
	vtClient           *vtclient.Client
	bscScanCrawler     *bscscan.Crawler
	onchainClient      *onchain.RestClient
	kyberswapBSCClient *kyberswapbsc.Kyberswapbsc
}

func NewStateAnalyzer(
	vtClient *vtclient.Client,
	bscScanCrawler *bscscan.Crawler,
	onchainClient *onchain.RestClient,
	kyberswapBSCClient *kyberswapbsc.Kyberswapbsc,
) *StateAnalyzer {
	return &StateAnalyzer{
		vtClient:           vtClient,
		bscScanCrawler:     bscScanCrawler,
		onchainClient:      onchainClient,
		kyberswapBSCClient: kyberswapBSCClient,
	}
}

func (a *StateAnalyzer) ExportSummary(ctx context.Context) ([]StateSummary, error) {
	states, err := a.getDoneStatesSortedByTime(ctx)
	if err != nil {
		return nil, err
	}

	var (
		summaries = make([]StateSummary, 0, len(states))
	)

	for _, s := range states {
		p1BaseFilled, p1QuoteFilled := a.summarizeCEXOrders(s.P1CEXOrders)
		p2BaseFilled, p2QuoteFilled := a.summarizeCEXOrders(s.P2CEXOrders)
		p2DEXSummary, err := a.summarizePart2DEX(ctx, s.P2DEXTxs)
		if err != nil {
			return nil, err
		}
		summaries = append(
			summaries,
			StateSummary{
				StateID:     s.StateID,
				CreatedTime: s.CreatedTime,
				P1Side:      s.Side,
				Part1: Part1Summary{
					P1OrderBaseFilled:  p1BaseFilled,
					P1OrderQuoteFilled: p1QuoteFilled,
				},
				Part2CEX: Part2CEXSummary{
					P2OrderBaseFilled:  p2BaseFilled,
					P2OrderQuoteFilled: p2QuoteFilled,
				},
				Part2DEX: p2DEXSummary,
			},
		)
	}

	return summaries, nil
}

func (a *StateAnalyzer) getDoneStatesSortedByTime(ctx context.Context) ([]vtclient.StateData, error) {
	states, err := a.vtClient.GetStates(ctx, true, true)
	if err != nil {
		return nil, err
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i].CreatedTime.Before(states[j].CreatedTime)
	})

	return states, nil
}

func (a *StateAnalyzer) summarizeCEXOrders(orders []vtclient.CEXOrderData) (baseFilled, quoteFilled float64) {
	for _, ord := range orders {
		baseFilled += ord.FilledBaseAmount
		quoteFilled += ord.FilledQuoteAmount
	}

	return baseFilled, quoteFilled
}

func (a *StateAnalyzer) summarizePart2DEX(ctx context.Context, txs []vtclient.DexTxData) (Part2DEXSummary, error) {
	var (
		totalGasFee                   float64
		totalAmountIn, totalAmountOut = new(big.Int), new(big.Int)
	)

	for _, tx := range txs {
		if tx.Status == "CANCELED" {
			continue
		}

		receipt, err := a.onchainClient.GetReceiptByTxHash(ctx, tx.TxHash)
		if err != nil {
			log.Printf("[ERROR] GetReceiptByTxHash tx: %+v", tx)
			return Part2DEXSummary{}, err
		}

		bnbPrice, err := a.bscScanCrawler.GetTxBNBPrice(tx.TxHash)
		if err != nil {
			return Part2DEXSummary{}, nil
		}
		totalGasFee += bnbPrice * float64(receipt.GasUsed)

		if receipt.Status == 0 {
			continue
		}

		swappedLog, err := kyberswapbsc.GetSwappedLog(a.kyberswapBSCClient, receipt.Logs)
		if err != nil {
			return Part2DEXSummary{}, nil
		}

		totalAmountIn.Add(totalAmountIn, swappedLog.SpentAmount)
		totalAmountOut.Add(totalAmountOut, swappedLog.ReturnAmount)
	}

	return Part2DEXSummary{
		P2TotalGasFee: totalGasFee,
		P2AmountIn:    totalAmountIn,
		P2AmountOut:   totalAmountOut,
	}, nil
}

type StateSummary struct {
	StateID     string
	CreatedTime time.Time
	P1Side      string

	Part1 Part1Summary

	Part2CEX Part2CEXSummary
	Part2DEX Part2DEXSummary
}

type Part1Summary struct {
	P1OrderBaseFilled  float64
	P1OrderQuoteFilled float64
}

type Part2CEXSummary struct {
	P2OrderBaseFilled  float64
	P2OrderQuoteFilled float64
}

type Part2DEXSummary struct {
	P2TotalGasFee float64
	P2AmountIn    *big.Int
	P2AmountOut   *big.Int
}
