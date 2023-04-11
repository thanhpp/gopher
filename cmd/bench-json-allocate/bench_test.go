package benchjsonallocate_test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	benchjsonallocate "github.com/thanhpp/gopher/cmd/bench-json-allocate"
)

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

func BenchmarkJSONiterAllocate(b *testing.B) {
	json := jsoniter.ConfigDefault
	for i := 0; i < b.N; i++ {
		var s benchjsonallocate.StaticExtra
		json.UnmarshalFromString(data, &s)
	}
}

func BenchmarkJSONAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s benchjsonallocate.StaticExtra
		json.Unmarshal([]byte(data), &s)
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s benchjsonallocate.StaticExtra
		easyjson.Unmarshal([]byte(data), &s)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-json-allocate
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkJSONiterAllocate
BenchmarkJSONiterAllocate-12    	  837956	      2240 ns/op	    1296 B/op	      19 allocs/op
BenchmarkJSONAllocate
BenchmarkJSONAllocate-12        	  281359	      3968 ns/op	     992 B/op	      14 allocs/op
BenchmarkEasyJSON
BenchmarkEasyJSON-12            	 1000000	      1266 ns/op	     728 B/op	      10 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-json-allocate	4.333s
*/
