package benchconcurrentmap_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/patrickmn/go-cache"
	"github.com/sourcegraph/conc"
	benchconcurrentmap "github.com/thanhpp/gopher/cmd/bench-concurrent-map"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var p = benchconcurrentmap.Pool{
	Address:    "0x06df3b2bbb68adc8b0e302443692037ed9f91b42",
	ReserveUsd: 0,
	SwapFee:    0.0004,
	Exchange:   "balancer",
	Type:       "balancer-stable",
	Timestamp:  13529165,
	Reserves: []string{"4362365955985",
		"4342743177527924936049411",
		"6921895060068041759669604",
		"4198113236810"},
	Tokens: []benchconcurrentmap.PoolToken{
		{
			Address: "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
			Weight:  250000000000000000,
		},
		{
			Address: "0x8f3cf7ad23cd3cadbd9735aff958023239c6a063",
			Weight:  250000000000000000,
		},
		{
			Address: "0xa3fa99a148fa48d14ed51d610c367c61876997f1",
			Weight:  250000000000000000,
		},
		{
			Address: "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
			Weight:  250000000000000000,
		},
	},
	Extra:       "{\"amplificationParameter\":{\"value\":60000,\"isUpdating\":false,\"precision\":1000}}",
	StaticExtra: "{\"vaultAddress\":\"0xba12222222228d8ba445958a75a0704d566bf2c8\",\"poolId\":\"0x06df3b2bbb68adc8b0e302443692037ed9f91b42000000000000000000000012\",\"tokenDecimals\":[6,18,18,6]}",
}

// NewPublicAddress testing only
func NewPublicAddress() common.Address {
	privateKey, _ := crypto.GenerateKey()
	pubAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	return pubAddr
}

func genData() ([]benchconcurrentmap.Pool, []string) {
	var (
		n      = 10_000
		modulo = 3
	)
	data := make([]benchconcurrentmap.Pool, n)
	searchKeys := []string{}
	for i := range data {
		data[i] = p.Clone()
		data[i].Address = NewPublicAddress().String()
		if i%modulo == 0 {
			searchKeys = append(searchKeys, data[i].Address)
		}
	}

	return data, searchKeys
}

func BenchmarkSimpleCache(b *testing.B) {
	data, search := genData()

	b.ResetTimer()

	for bN := 0; bN < b.N; bN++ {
		wg := conc.WaitGroup{}
		sc := benchconcurrentmap.NewSimpleCache()

		wg.Go(func() {
			for i := range data[:len(data)/2] {
				sc.Set(data[i])
			}
		})
		wg.Go(func() {
			for i := range data[len(data)/2:] {
				sc.Set(data[len(data)/2+i])
			}
		})
		wg.Go(func() {
			sc.Get(search...)
		})

		wg.Wait()
	}
}

func BenchmarkConcMap(b *testing.B) {
	data, search := genData()

	b.ResetTimer()

	for bN := 0; bN < b.N; bN++ {
		wg := conc.WaitGroup{}
		m := cmap.New[benchconcurrentmap.Pool]()

		wg.Go(func() {
			for i := range data[:len(data)/2] {
				m.Set(data[i].Address, data[i])
			}
		})
		wg.Go(func() {
			for i := range data[len(data)/2:] {
				m.Set(data[i].Address, data[len(data)/2+i])
			}
		})
		wg.Go(func() {
			for i := range search {
				m.Get(search[i])
			}
		})

		wg.Wait()
	}
}

func BenchmarkGoCache(b *testing.B) {
	data, search := genData()

	b.ResetTimer()

	for bN := 0; bN < b.N; bN++ {
		wg := conc.WaitGroup{}
		m := cache.New(cache.NoExpiration, cache.NoExpiration)

		wg.Go(func() {
			for i := range data[:len(data)/2] {
				m.Set(data[i].Address, data[i], cache.NoExpiration)
			}
		})
		wg.Go(func() {
			for i := range data[len(data)/2:] {
				m.Set(data[i].Address, data[len(data)/2+i], cache.NoExpiration)
			}
		})
		wg.Go(func() {
			for i := range search {
				m.Get(search[i])
			}
		})

		wg.Wait()
	}
}
