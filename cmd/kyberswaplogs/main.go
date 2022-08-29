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
	bscURL                   = "https://bsc-dataseed.binance.org"
	txHash                   = "0xcd76d6acce33c3b7903413eceb86548d0457a55811f48056da2f13f5c26d1245"
	kyberswapBSCSwappedTopic = "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8"
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
		if len(receipt.Logs[i].Topics) == 0 {
			continue
		}

		for j := range receipt.Logs[i].Topics {
			if receipt.Logs[i].Topics[j].String() == kyberswapBSCSwappedTopic {
				result, err := kyberswapBSC.ParseSwapped(*receipt.Logs[i])
				if err != nil {
					log.Println("[SKIP ERROR] kyberswapbsc parse swapped error:", err)
					continue
				}

				log.Printf("reading tx receipt: %+v", receipt.Logs[i])
				log.Printf("kyberswapbsc parse swapped: %+v", result)
			}
		}
	}
}
