package onchain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thanhpp/gopher/pkg/blockchain/chains"
)

type RestClient struct {
	chain chains.Chain
	ethC  *ethclient.Client
}

func NewRestClient(chain chains.Chain, url string) (*RestClient, error) {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("ethclient dial error: %w", err)
	}

	c := &RestClient{
		chain: chain,
		ethC:  ethClient,
	}

	return c, nil
}

func (r *RestClient) Chain() chains.Chain {
	return r.chain
}

func (r *RestClient) Client() *ethclient.Client {
	return r.ethC
}

func (r *RestClient) GetReceiptByTxHash(ctx context.Context, txHash string) (*types.Receipt, error) {
	receipt, err := r.ethC.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return nil, fmt.Errorf("get tx receipt error: %w, txHash: %s", err, txHash)
	}

	return receipt, nil
}
