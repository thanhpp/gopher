package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sourcegraph/conc"
	"github.com/thanhpp/gopher/pkg/kyberswapcontract"
)

const (
	bscRPC   = "https://bsc-dataseed2.binance.org"
	testAddr = "0xB81eFe13BF37F7045E923d005D3073Cd3917F30e"
)

var (
	ctx      = context.Background()
	from     = common.HexToAddress(testAddr)
	ksAbi, _ = abi.JSON(bytes.NewBufferString(kyberswapcontract.KyberswapcontractABI))
)

func main() {
	ethClient, err := ethclient.Dial(bscRPC)
	if err != nil {
		log.Fatalf("dial ethclient error: %v", err)
	}
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		log.Fatalf("get chainID error: %v", err)
	}
	log.Printf("chainID: %s", chainID.String())

	ticker := time.NewTicker(time.Minute)
	for ; ; <-ticker.C {
		var (
			wg                            conc.WaitGroup
			gasEstimated, gasCallContract uint64
		)
		data, router, err := GetKSData(testAddr)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := ethClient.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatalf("suggest gasPrice error: %v", err)
		}

		wg.Go(func() {
			gas, err := GetGasLimit(ctx, ethClient, from, common.HexToAddress(router), gasPrice, data)
			if err != nil {
				log.Fatalf("get gas limit error: %v", err)
			}
			gasEstimated = gas
		})
		wg.Go(func() {
			gas, err := GetGasByCallContract(ctx, ethClient, from, common.HexToAddress(router), gasPrice, data)
			if err != nil {
				log.Fatalf("get gas no send error: %v", err)
			}
			gasCallContract = gas
		})
		wg.Wait()

		log.Println("estimated", gasEstimated, "callContract", gasCallContract,
			"diff", (gasEstimated - gasCallContract))
	}
}

func GetKey(privateKeyFile, password string) (*keystore.Key, error) {
	data, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("read private key file error: %w", err)
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, fmt.Errorf("decrypt private key error: %w", err)
	}

	return key, nil
}

func GetKSData(to string) (data []byte, router string, err error) {
	url := "https://aggregator-api.kyberswap.com/bsc/route/encode?tokenIn=0x55d398326f99059fF775485246999027B3197955&tokenOut=0x715D400F88C167884bbCc41C5FeA407ed4D2f8A0&amountIn=1500000000000000000000&saveGas=0&gasInclude=0&slippageTolerance=50&to=" // nolint: lll
	sendURL := fmt.Sprintf("%s%s", url, to)

	resp, err := http.Get(sendURL)
	if err != nil {
		return nil, "", fmt.Errorf("get ks get request error: %w", err)
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("read response data error: %w", err)
	}

	ksResp := new(GetSwapInfoResponse)
	if err := json.Unmarshal(bodyData, ksResp); err != nil {
		return nil, "", err
	}

	return hexutil.MustDecode(ksResp.EncodedSwapData), ksResp.RouterAddress, nil
}

type GetSwapInfoResponse struct {
	EncodedSwapData string `json:"encodedSwapData"`
	RouterAddress   string `json:"routerAddress"`
}

func GetGasLimit(
	ctx context.Context, ethClient *ethclient.Client, from, routerAddr common.Address, gasPrice *big.Int,
	data []byte,
) (uint64, error) {
	callMsg := ethereum.CallMsg{
		From:     from,
		To:       &routerAddr,
		GasPrice: gasPrice,
		Data:     data,
		Value:    big.NewInt(0),
	}

	gasLimit, err := ethClient.EstimateGas(ctx, callMsg)
	if err != nil {
		log.Printf("[ERROR] EstimateGas callMsg: %+v", callMsg)
		return 0, fmt.Errorf("estimate gas error: %w", err)
	}

	return gasLimit, nil
}

func GetGasByCallContract(
	ctx context.Context, ethClient *ethclient.Client, from,
	routerAddr common.Address, gasPrice *big.Int, data []byte,
) (uint64, error) {
	callMsg := ethereum.CallMsg{
		From:     from,
		To:       &routerAddr,
		GasPrice: gasPrice,
		Data:     data,
		Value:    big.NewInt(0),
	}

	callData, err := ethClient.CallContract(ctx, callMsg, nil)
	if err != nil {
		return 0, fmt.Errorf("call contract error: %w", err)
	}

	outputs, err := ksAbi.Unpack("swap", callData)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	gas, ok := outputs[1].(*big.Int)
	if !ok {
		return 0, fmt.Errorf("cast gas error, outputs: %v, outputs[1]: %T", outputs, outputs[1])
	}

	return gas.Uint64(), nil
}
