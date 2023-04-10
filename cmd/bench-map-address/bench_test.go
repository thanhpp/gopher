package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// NewPublicAddress testing only
func NewPublicAddress() common.Address {
	privateKey, _ := crypto.GenerateKey()
	pubAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	return pubAddr
}

func genKey() ([]common.Address, []common.Address) {
	var (
		n       = 10_000
		modulo  = 8
		result  = make([]common.Address, 0, n)
		getKeys []common.Address
	)

	for i := 0; i < n; i++ {
		addr := NewPublicAddress()
		result = append(result, addr)
		if i%modulo == 0 {
			getKeys = append(getKeys, addr)
		}
	}

	return result, getKeys
}

func BenchmarkMapAddress(b *testing.B) {
	keys, getKeys := genKey()
	m := make(map[common.Address]string, len(keys))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			m[k] = "abc"
		}

		for _, k := range getKeys {
			if _, ok := m[k]; !ok {
				b.Logf("key not found: %+v", k)
			}
		}
	}
}

func BenchmarkMapString(b *testing.B) {
	keys, getKeys := genKey()
	m := make(map[string]string, len(keys))

	strKeys := make([]string, 0, len(keys))
	for i := range keys {
		strKeys = append(strKeys, keys[i].String())
	}

	strGetKeys := make([]string, 0, len(keys))
	for i := range getKeys {
		strGetKeys = append(strGetKeys, getKeys[i].String())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, k := range strKeys {
			m[k] = "abc"
		}

		for _, k := range strGetKeys {
			if _, ok := m[k]; !ok {
				b.Logf("key not found: %+v", k)
			}
		}
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-map-address
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkMapAddress
BenchmarkMapAddress-12    	    3733	    320354 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapString
BenchmarkMapString-12     	    4303	    281193 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-map-address	5.473s
*/
