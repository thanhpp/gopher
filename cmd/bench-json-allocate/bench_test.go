package main

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type StaticExtra struct {
	PoolId           string   `json:"poolId"`
	LpToken          string   `json:"lpToken"`
	Type             string   `json:"type"`
	Tokens           []string `json:"tokens"`
	DodoV1SellHelper string   `json:"dodoV1SellHelper"`
}

var (
	data = `{
	"poolId": "poolIdpoolIdpoolIdpoolId",
	"lpToken": "lpTokenlpTokenlpTokenlpTokenlpToken",
	"type": "typetypetypetypetypetypetype",
	"tokens": [
		"tokenstokenstokenstokenstokenstokens",
		"tokenstokenstokenstokens",
		"tokenstokens",
		"tokenstokenstokenstokens"
	],
	"dodoV1SellHelper": "dodoV1SellHelperdodoV1SellHelperdodoV1SellHelperdodoV1SellHelper"
}`
)

func BenchmarkJSONiterNonAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s StaticExtra
		jsoniter.ConfigDefault.UnmarshalFromString(data, &s)
	}
}

func BenchmarkJSONiterAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := new(StaticExtra)
		jsoniter.ConfigDefault.UnmarshalFromString(data, s)
	}
}

func BenchmarkJSONNonAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s StaticExtra
		json.Unmarshal([]byte(data), &s)
	}
}

func BenchmarkJSONAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := new(StaticExtra)
		json.Unmarshal([]byte(data), s)
	}
}

/*
1.18

goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-json-allocate
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkJSONiterNonAllocate
BenchmarkJSONiterNonAllocate-12    	 1168766	      1042 ns/op	     920 B/op	      18 allocs/op
BenchmarkJSONiterAllocate
BenchmarkJSONiterAllocate-12       	 1000000	      1011 ns/op	     920 B/op	      18 allocs/op
BenchmarkJSONNonAllocate
BenchmarkJSONNonAllocate-12        	  383960	      3174 ns/op	    1080 B/op	      17 allocs/op
BenchmarkJSONAllocate
BenchmarkJSONAllocate-12           	  377635	      3063 ns/op	    1080 B/op	      17 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-json-allocate	7.563s
*/

/*
1.20

goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-json-allocate
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkJSONiterNonAllocate
BenchmarkJSONiterNonAllocate-12    	 1212568	       988.3 ns/op	     920 B/op	      18 allocs/op
BenchmarkJSONiterAllocate
BenchmarkJSONiterAllocate-12       	 1208973	       999.4 ns/op	     920 B/op	      18 allocs/op
BenchmarkJSONNonAllocate
BenchmarkJSONNonAllocate-12        	  344893	      2979 ns/op	    1080 B/op	      17 allocs/op
BenchmarkJSONAllocate
BenchmarkJSONAllocate-12           	  376743	      3074 ns/op	    1080 B/op	      17 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-json-allocate	6.670s
*/
