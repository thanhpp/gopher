package bscscan_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanhpp/gopher/pkg/crawlers/bscscan"
)

func TestGetTxBNBPrice(t *testing.T) {
	var (
		c      = bscscan.New()
		txHash = "0x5415817fbc3aff6db42b33845cf127acbb8e4b895b4a57f8da846d1944562a8d"
	)

	bnbPrice, err := c.GetTxBNBPrice(txHash)
	assert.NoError(t, err)

	t.Log("BNB price", bnbPrice)
}
