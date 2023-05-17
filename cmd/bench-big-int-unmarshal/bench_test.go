package benchbigintunmarshal_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type A struct {
	B *big.Int `json:"b"`
}

type A1 struct {
	B string `json:"b"`
}

var data = `{
	"b": 120938120487198236129837129847359843091851098236479326419875436732154187345182736451823764
}
`

func BenchmarkUnmarshalBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a A
		assert.NoError(b, json.Unmarshal([]byte(data), &a))
	}
}

func TestUnmarshalBigInt(t *testing.T) {
	var a A
	assert.NoError(t, json.Unmarshal([]byte(data), &a))
	t.Log(a.B.String())
}
