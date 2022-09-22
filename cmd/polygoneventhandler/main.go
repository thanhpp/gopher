package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/thanhpp/gopher/cmd/polygoneventhandler/smartcontracts"
	"go.uber.org/zap"
)

const (
	envFile        = "./.env"
	nscAddressEnv  = "NSC_ADDRESS"
	mumbaiWsURLEnv = "MUMBAI_WEBSOCKET"
)

var (
	nscAddress common.Address
	// transferKeccak256 Transfer(address,address,uint256)
	transferKeccak256 = common.HexToHash("ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	// approvalKeccak256 Approval(address,address,uint256)
	approvalKeccak256 = common.HexToHash("8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
	mumbaiWsURL       string
	smInstance        *smartcontracts.Smartcontracts
)

func main() {
	setZap()
	loadEnv()

	ethClient, err := ethclient.Dial(mumbaiWsURL)
	if err != nil {
		log.Fatal("main ethclient.Dial error", err)
	}

	smInstance, err = smartcontracts.NewSmartcontracts(nscAddress, ethClient)
	if err != nil || smInstance == nil {
		log.Fatal("smartcontracts.NewSmartcontracts error", err)
	}

	if err := startSub(context.Background(), ethClient); err != nil {
		log.Fatalf("startSub error: %v", err)
		return
	}
}

func loadEnv() {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("load env error", err)
	}

	nscAddress = common.HexToAddress(os.Getenv(nscAddressEnv))
	mumbaiWsURL = os.Getenv(mumbaiWsURLEnv)

	zap.S().Debugw("loadEnv", "nscAddress", nscAddress, "mumbaiWsURL", mumbaiWsURL)
}

func setZap() {
	zlogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(zlogger)
}

func startSub(ctx context.Context, ethClient *ethclient.Client) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{nscAddress},
	}
	logsC := make(chan types.Log)

	sub, err := ethClient.SubscribeFilterLogs(ctx, query, logsC)
	if err != nil {
		return fmt.Errorf("ethClient.SubscribeFilterLogs error: %w", err)
	}

	handleSubcription(sub, logsC)

	return nil
}

func handleSubcription(sub ethereum.Subscription, logsC chan types.Log) {
	log.Println("handleSubcription starting...")
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("handleSubcription - sub.Err error: %v", err)
		case lg := <-logsC:
			handleLogs(lg)
		}
	}
}

func handleLogs(lg types.Log) {
	if len(lg.Topics) == 0 {
		log.Printf("[DEBUG] empty topic logs - %+v", lg)
		return
	}

	switch lg.Topics[0] {
	case transferKeccak256:
		transferInfo, err := smInstance.ParseTransfer(lg)
		if err != nil {
			log.Println("[DEBUG] Parse Transfer err", err, "log", lg)
			return
		}
		log.Printf("[DEBUG] transfer info: %+v", transferInfo)

	case approvalKeccak256:
		approvalInfo, err := smInstance.ParseApproval(lg)
		if err != nil {
			log.Println("[DEBUG] Parse Approval err", err, "log", lg)
			return
		}
		log.Printf("[DEBUG] Approvel info: %+v", approvalInfo)

	default:
		log.Println("unsupported topic", lg.Topics[0], lg)
	}
}

/*
mint(to 0x532da84493eb17968654cf3057af8afb634223e3, _tokenIds 10) => success

2022/09/21 23:02:56 handleSubcription - new log: {Address:0xF3797827f8a4e15f726D057579C135F001176615 Topics:[0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef 0x0000000000000000000000000000000000000000000000000000000000000000 0x000000000000000000000000532da84493eb17968654cf3057af8afb634223e3 0x000000000000000000000000000000000000000000000000000000000000000a] Data:[] BlockNumber:28221693 TxHash:0x6f0cee62e2b46dfab6cc3de2f1dd534f9e283e8f80ffba40c4804c2d2308f746 TxIndex:31 BlockHash:0x97a0d81e4ad61af84c211780c112f7c7e52f50474f42007f071bdb42d8e66fbc Index:213 Removed:false}
*/ // nolint

/*
mint 886260330
2022/09/22 10:48:13 handleSubcription - new log: {Address:0xF3797827f8a4e15f726D057579C135F001176615 Topics:[0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925 0x000000000000000000000000532da84493eb17968654cf3057af8afb634223e3 0x0000000000000000000000000000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000000000000000000034d3426a] Data:[] BlockNumber:28230130 TxHash:0x955762e6aaedbcc7d2ec3699002a64d24c37b6b766e9e67fc92cb3d91185a637 TxIndex:25 BlockHash:0x210403d5d26be2cd8b78e7e50a2a6c78d484c9135ef37d9c12fb340ae8d17622 Index:112 Removed:false}

burn 886260330
2022/09/22 10:48:13 handleSubcription - new log: {Address:0xF3797827f8a4e15f726D057579C135F001176615 Topics:[0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef 0x000000000000000000000000532da84493eb17968654cf3057af8afb634223e3 0x0000000000000000000000000000000000000000000000000000000000000000 0x0000000000000000000000000000000000000000000000000000000034d3426a] Data:[] BlockNumber:28230130 TxHash:0x955762e6aaedbcc7d2ec3699002a64d24c37b6b766e9e67fc92cb3d91185a637 TxIndex:25 BlockHash:0x210403d5d26be2cd8b78e7e50a2a6c78d484c9135ef37d9c12fb340ae8d17622 Index:113 Removed:false}
*/ // nolint
