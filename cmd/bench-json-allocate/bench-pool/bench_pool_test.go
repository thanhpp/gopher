package benchpool_test

import (
	json "encoding/json"
	"testing"

	gjson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	easyjson "github.com/mailru/easyjson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	benchpool "github.com/thanhpp/gopher/cmd/bench-json-allocate/bench-pool"
)

// nolint: lll
var data = "{\"swapFee\":100,\"exchange\":\"uniswapv3\",\"type\":\"uniswapv3\",\"timestamp\":1681201870,\"reserves\":[\"5473051155136514777\",\"1962858960281861122\"],\"tokens\":[{\"address\":\"0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270\",\"name\":\"Wrapped Matic\",\"symbol\":\"WMATIC\",\"decimals\":18,\"weight\":50,\"swappable\":true},{\"address\":\"0x0fa9ab8ff0e539095aaf44744a1f37d94e4fb1d1\",\"name\":\"USDC\",\"symbol\":\"USDC\",\"decimals\":18,\"weight\":50,\"swappable\":true}],\"extra\":\"{\\\"liquidity\\\":2963051155136514777,\\\"sqrtPriceX96\\\":52484314497306184951177414416,\\\"tick\\\":-8237,\\\"ticks\\\":[{\\\"index\\\":0,\\\"liquidityGross\\\":3414550863629896216,\\\"liquidityNet\\\":3414550863629896216},{\\\"index\\\":6931,\\\"liquidityGross\\\":3414550863629896216,\\\"liquidityNet\\\":-3414550863629896216},{\\\"index\\\":887272,\\\"liquidityGross\\\":2963051155136514777,\\\"liquidityNet\\\":-2963051155136514777},{\\\"index\\\":-887272,\\\"liquidityGross\\\":2963051155136514777,\\\"liquidityNet\\\":2963051155136514777}]}\",\"staticExtra\":\"{\\\"poolId\\\":\\\"0x87a9f01bfdc4fcd7ce223bd667871a7c4d9fd3b0\\\"}\"}"

func TestUnmarshal(t *testing.T) {
	var p1 benchpool.Pool
	require.NoError(t, json.Unmarshal([]byte(data), &p1))

	var p2 benchpool.Pool
	jsonit := jsoniter.ConfigDefault
	require.NoError(t, jsonit.UnmarshalFromString(data, &p2))

	var p3 benchpool.Pool
	require.NoError(t, easyjson.Unmarshal([]byte(data), &p3))

	var p4 benchpool.Pool
	require.NoError(t, gjson.Unmarshal([]byte(data), &p4))

	assert.Equal(t, p1, p2)
	assert.Equal(t, p1, p3)
	assert.Equal(t, p1, p4)
}

func BenchmarkJSONiterAllocate(b *testing.B) {
	json := jsoniter.ConfigDefault
	for i := 0; i < b.N; i++ {
		var p benchpool.Pool
		json.UnmarshalFromString(data, &p)
	}
}

func BenchmarkJSONAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p benchpool.Pool
		json.Unmarshal([]byte(data), &p)
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p benchpool.Pool
		easyjson.Unmarshal([]byte(data), &p)
	}
}

func BenchmarkGoJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var p benchpool.Pool
		gjson.Unmarshal([]byte(data), &p)
	}
}
