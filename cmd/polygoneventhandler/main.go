package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	envFile        = "./.env"
	nscAddressEnv  = "NSC_ADDRESS"
	mumbaiWsURLEnv = "MUMBAI_WEBSOCKET"
)

var (
	nscAddress  common.Address
	mumbaiWsURL string
)

func main() {
	setZap()
	loadEnv()

	_, err := ethclient.Dial(mumbaiWsURL)
	if err != nil {
		log.Fatal("main ethclient.Dial error", err)
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
