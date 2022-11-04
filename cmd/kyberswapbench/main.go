package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

// https://docs.kyberswap.com/Aggregator/aggregator-api

type Info struct {
	Amount     float64
	PriceWoGas float64
	PriceWGas  float64
	Gas        float64
	Diff       float64
}

func main() { // nolint
	var (
		min   = 1_500
		max   = 3_000
		step  = 250
		wg    sync.WaitGroup
		m     = make(map[int64]*KyberSwapResp)
		mlock sync.Mutex
	)

	f, err := os.Create(fmt.Sprintf("%s.txt", time.Now().String()))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		f.WriteString("---------------\n")
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
			bestAmount     int64
			bestPrice      float64
			bestDiffAmount int64
			bestDiff       float64 = math.MaxFloat64
			info                   = make([]Info, 0, len(m))
		)
		for k := min; k <= max; k += step {
			v := m[int64(k)]
			in, _ := new(big.Int).SetString(v.InputAmount, 10)
			out, _ := new(big.Int).SetString(v.OutputAmount, 10)
			priceWoGas := TokenAmountToFloat(out, 18) / TokenAmountToFloat(in, 18)
			priceWGas := priceWoGas * (1 - v.GasUsd/v.AmountOutUsd)
			diff := 0.001 + 0.002 + v.GasUsd/v.AmountOutUsd
			// f.WriteString(fmt.Sprintf("%d: %f\tgasUsd: %f\t diff: %f\n", k, price, v.GasUsd, diff))

			info = append(info, Info{
				Amount:     float64(k),
				PriceWGas:  priceWGas,
				PriceWoGas: priceWoGas,
				Diff:       diff,
				Gas:        v.GasUsd,
			})

			if priceWGas > bestPrice {
				bestAmount = int64(k)
				bestPrice = priceWGas
			}

			if diff < bestDiff {
				bestDiff = diff
				bestDiffAmount = int64(k)
			}
		}

		log.Printf("writing to file\n")
		sort.Slice(info, func(i, j int) bool {
			return info[i].PriceWGas < info[j].PriceWGas
		})
		for i := range info {
			f.WriteString(fmt.Sprintf("Amount: %5.5f, PriceWGas: %5.5f, PriceWoGas: %5.5f, Diff: %5.5f, Gas: %5.5f\n",
				info[i].Amount, info[i].PriceWGas, info[i].PriceWoGas, info[i].Diff, info[i].Gas))
		}

		f.WriteString(fmt.Sprintf("BEST: %d: %f\n", bestAmount, bestPrice))
		f.WriteString(fmt.Sprintf("BEST DIFF: %d: %f\n\n", bestDiffAmount, bestDiff))
	}
}

func getKSRate(amount int64) (*KyberSwapResp, error) {
	// url := fmt.Sprintf("https://aggregator-api.kyberswap.com/ethereum/route/encode?tokenIn=0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202&tokenOut=0xdAC17F958D2ee523a2206206994597C13D831ec7&amountIn=%s&to=0x000000000000000000000000000000000000dEaD", IntToTokenAmount(amount, 18).String()) // nolint: lll
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
		log.Printf("%+v", resp.Header)
		return nil, errors.New("not 200")
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	ksResp := new(KyberSwapResp)

	log.Println(string(data))

	if err := json.Unmarshal(data, ksResp); err != nil {
		return nil, err
	}

	log.Println(ksResp)

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
