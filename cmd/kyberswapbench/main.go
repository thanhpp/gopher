package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"
)

// https://docs.kyberswap.com/Aggregator/aggregator-api

func main() {
	var (
		min   = 600
		max   = 1_000
		step  = 200
		wg    sync.WaitGroup
		m     = make(map[int64]*KyberSwapResp)
		mlock sync.Mutex
	)

	f, err := os.Create(fmt.Sprintf("%s.txt", time.Now().String()))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		for i := min; i <= max; i += step {
			wg.Add(1)
			go func(amount int64) {
				log.Printf("start call KS: amount %d\n", amount)
				defer log.Printf("done call KS: amount %d\n", amount)
				resp, err := getKSRate(amount)
				if err != nil {
					panic(err)
				}
				mlock.Lock()
				m[amount] = resp
				mlock.Unlock()
				wg.Done()
			}(int64(i))
		}
		wg.Wait()

		log.Printf("comparing price\n")
		var (
			bestAmount int64
			bestPrice  float64
		)
		for k, v := range m {
			in, _ := new(big.Int).SetString(v.InputAmount, 10)
			out, _ := new(big.Int).SetString(v.OutputAmount, 10)
			price := TokenAmountToFloat(out, 18) / TokenAmountToFloat(in, 18)
			f.WriteString(fmt.Sprintf("%d: %f\tgasUsd: %f\n", k, price, v.GasUsd))

			if price > bestPrice {
				bestAmount = k
				bestPrice = price
			}
		}

		log.Printf("writing to file\n")
		f.WriteString(fmt.Sprintf("BEST: %d: %f\n\n", bestAmount, bestPrice))
	}
}

func getKSRate(amount int64) (*KyberSwapResp, error) {
	url := fmt.Sprintf("https://aggregator-api.kyberswap.com/bsc/route/encode?tokenIn=0xfe56d5892BDffC7BF58f2E84BE1b2C32D21C308b&tokenOut=0x55d398326f99059fF775485246999027B3197955&amountIn=%s&to=0x000000000000000000000000000000000000dEaD", IntToTokenAmount(amount, 18).String()) // nolint: lll
	log.Printf("[DEBUG] %s\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not 200")
	}
	defer resp.Body.Close()

	ksResp := new(KyberSwapResp)

	if err := json.NewDecoder(resp.Body).Decode(ksResp); err != nil {
		return nil, err
	}

	return ksResp, nil
}

func IntToTokenAmount(amount int64, decimals int64) *big.Int {
	weiFloat := big.NewInt(amount)
	decimalsBig := Exp10(decimals)
	amountBig := new(big.Int).Mul(weiFloat, decimalsBig)

	return amountBig
}

func Exp10(n int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(n), nil)
}

func TokenAmountToFloat(amount *big.Int, decimals int64) float64 {
	amountFloat := big.NewFloat(0).SetInt(amount)
	amountFloat.Quo(amountFloat, big.NewFloat(0).SetInt(Exp10(decimals)))

	output, _ := amountFloat.Float64()

	return output
}

type KyberSwapResp struct {
	InputAmount  string  `json:"inputAmount"`
	OutputAmount string  `json:"outputAmount"`
	TotalGas     int     `json:"totalGas"`
	GasPriceGwei string  `json:"gasPriceGwei"`
	GasUsd       float64 `json:"gasUsd"`
	AmountInUsd  float64 `json:"amountInUsd"`
	AmountOutUsd float64 `json:"amountOutUsd"`
	ReceivedUsd  float64 `json:"receivedUsd"`
	Swaps        [][]struct {
		Pool              string `json:"pool"`
		TokenIn           string `json:"tokenIn"`
		TokenOut          string `json:"tokenOut"`
		SwapAmount        string `json:"swapAmount"`
		AmountOut         string `json:"amountOut"`
		LimitReturnAmount string `json:"limitReturnAmount"`
		MaxPrice          string `json:"maxPrice"`
		Exchange          string `json:"exchange"`
		PoolLength        int    `json:"poolLength"`
		PoolType          string `json:"poolType"`
	} `json:"swaps"`
	EncodedSwapData string `json:"encodedSwapData"`
	RouterAddress   string `json:"routerAddress"`
}
