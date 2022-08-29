package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/thanhpp/gopher/internal/vtnoti/domain/service"
	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/blockchain/chains"
	"github.com/thanhpp/gopher/pkg/blockchain/contracts/kyberswapbsc"
	"github.com/thanhpp/gopher/pkg/blockchain/onchain"
	"github.com/thanhpp/gopher/pkg/crawlers/bscscan"
)

const (
	vtClientEnvFile      = "../../secrets/vtclient.env"
	virtualTakerBaseURL  = "VIRTUAL_TAKER_BASE_URL"
	virtualTakerUsername = "VIRTUAL_TAKER_USERNAME"
	virtualTakerPassword = "VIRTUAL_TAKER_PASSWORD"
)

const (
	bscURL = "https://bsc-dataseed.binance.org"
)

func main() {
	if err := godotenv.Load(vtClientEnvFile); err != nil {
		log.Println("[ERROR] load env", err)
		return
	}

	vtClient := vtclient.NewClient(
		os.Getenv(virtualTakerBaseURL), os.Getenv(virtualTakerUsername), os.Getenv(virtualTakerPassword),
	)

	bscOnchainClient, err := onchain.NewRestClient(chains.BSC, bscURL)
	if err != nil {
		log.Println("[ERROR] new bsc client", err)
		return
	}

	kyberswapbscClient, err := kyberswapbsc.NewKyberswapbsc(
		common.HexToAddress(kyberswapbsc.ContractAddress),
		bscOnchainClient.Client(),
	)
	if err != nil {
		log.Println("[ERROR] new kyberswap bsc", err)
		return
	}

	s := service.NewStateAnalyzer(
		vtClient,
		bscscan.New(),
		bscOnchainClient,
		kyberswapbscClient,
	)

	summaries, err := s.ExportSummary(context.Background())
	if err != nil {
		log.Println("[ERROR] export summary", err)
		return
	}

	for i := range summaries {
		log.Printf("%+v", summaries[i])
	}

	if err := writeCSV(summaries); err != nil {
		log.Println("[ERROR] write csv summary", err)
		return
	}
}

func writeCSV(summaries []service.StateSummary) error {
	f, err := os.Create(fmt.Sprintf("state-summary-%d.csv", time.Now().UnixNano()))
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	cols := []string{
		"state_id", "created_time",
		"p1_state_side", "p1_base_filled", "p1_quote_filled",
		"p2_base_filled", "p2_quote_filled",
		"p2_total_tx_fee", "p2_amount_in", "p2_amount_out",
	}
	if err := w.Write(cols); err != nil {
		return err
	}

	for _, s := range summaries {
		if err := w.Write([]string{
			s.StateID, s.CreatedTime.String(),
			s.P1Side, float64ToString(s.Part1.P1OrderBaseFilled), float64ToString(s.Part1.P1OrderQuoteFilled),
			float64ToString(s.Part2CEX.P2OrderBaseFilled), float64ToString(s.Part2CEX.P2OrderQuoteFilled),
			float64ToString(s.Part2DEX.P2TotalGasFee), s.Part2DEX.P2AmountIn.String(), s.Part2DEX.P2AmountOut.String(),
		},
		); err != nil {
			return err
		}
	}

	return nil
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 5, 64)
}
