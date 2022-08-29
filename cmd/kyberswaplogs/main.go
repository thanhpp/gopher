package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/thanhpp/gopher/pkg/blockchain/chains"
	"github.com/thanhpp/gopher/pkg/blockchain/contracts/kyberswapbsc"
	"github.com/thanhpp/gopher/pkg/blockchain/onchain"
)

var (
	bscURL = "https://bsc-dataseed.binance.org"
	txHash = "0xcd76d6acce33c3b7903413eceb86548d0457a55811f48056da2f13f5c26d1245"
)

func main() {
	ctx := context.Background()

	onchainClient, err := onchain.NewRestClient(chains.BSC, bscURL)
	if err != nil {
		log.Println(err)
		return
	}

	receipt, err := onchainClient.GetReceiptByTxHash(ctx, txHash)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("tx success:", receipt.Status == 1)

	kyberswapBSC, err := kyberswapbsc.NewKyberswapbsc(
		common.HexToAddress(kyberswapbsc.ContractAddress), onchainClient.Client())
	if err != nil {
		log.Println("new kyberswap bsc error:", err)
	}

	for i := range receipt.Logs {
		result, err := kyberswapBSC.ParseSwapped(*receipt.Logs[i])
		if err != nil {
			log.Println("[SKIP ERROR] kyberswapbsc parse swapped error:", err)
			continue
		}

		log.Printf("kyberswapbsc parse swapped: %+v", result)
	}
}
