// nolint:lll
package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	_ "embed"
)

//go:embed data.txt
var encodedData string

var (
	signatureEncoded = "0xa3b2a93b36a43917a1eccdc0a774fe4d4f2bfb5ce2e7556c8cc4f44e466e77a07791ab6bb05a1a68611e3b50480625c208f5e309fd469d357812eb8f43defba41b"
)

func main() {
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"Request": {
				{Name: "from", Type: "address"},
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
				{Name: "data", Type: "bytes"},
			},
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Request",
		Domain: apitypes.TypedDataDomain{
			Name:              "CexDexFund",
			Version:           "2.0",
			ChainId:           math.NewHexOrDecimal256(137),
			VerifyingContract: "0x2c46013298701F455E43fB7b3e3FF3129FD31226",
		},
		Message: apitypes.TypedDataMessage{},
	}

	typedDataMsg := apitypes.TypedDataMessage{
		"from":     "0xb0180eb030ac7322a1040feeb7cdd262985b6512",
		"to":       "0x6131b5fae19ea4f9d964eac0408e4408b66337b5",
		"value":    "0",
		"deadline": "1684294283",
		"data":     hexutil.MustDecode(encodedData),
	}

	// encode typed data message & hash
	newTypedData := typedData
	newTypedData.Message = typedDataMsg
	hash, _, err := apitypes.TypedDataAndHash(newTypedData)
	if err != nil {
		log.Panicf("hash typed data error: %v", err)
	}

	decoded := hexutil.MustDecode(signatureEncoded)
	if len(decoded) > 64 {
		decoded[64] -= 27
		// Transform V from 0/1 to 27/28 according to the yellow paper:https://eips.ethereum.org/EIPS/eip-712
	}
	sigPublicKey, err := crypto.SigToPub(hash, decoded)
	if err != nil {
		log.Panicf("Ecrecoverer: %v", err)
	}

	log.Println(crypto.PubkeyToAddress(*sigPublicKey)) // 0x36fe663005aC9f047b3cfa4f250BadDB16644A89
}
