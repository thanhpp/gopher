package main

import (
	"math/rand"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

type A struct {
	I int
}

var (
	n = 10_000
)

func genData() []A {
	sliceA := make([]A, n)

	for i := 0; i < n; i++ {
		sliceA = append(sliceA, A{
			I: rand.Intn(n),
		})
	}

	return sliceA
}

func BenchmarkNonGenericSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := genData()
		sort.SliceStable(data, func(i, j int) bool {
			return data[i].I < data[j].I
		})
	}
}

func BenchmarkGenericSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := genData()
		slices.SortStableFunc(data, func(a, b A) bool {
			return a.I < b.I
		})
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-generic-sort
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNonGenericSort
BenchmarkNonGenericSort-12    	     543	   3031778 ns/op	  507961 B/op	       6 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12       	     708	   2603175 ns/op	  507905 B/op	       4 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-generic-sort	3.888s

goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-generic-sort
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNonGenericSort
BenchmarkNonGenericSort-12    	     558	   2773717 ns/op	  507966 B/op	       6 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12       	     505	   2399293 ns/op	  507905 B/op	       4 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-generic-sort	4.201s

BenchmarkNonGenericSort
BenchmarkNonGenericSort-12    	     510	   2822796 ns/op	  507962 B/op	       6 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12       	     469	   2319975 ns/op	  507905 B/op	       4 allocs/op
PASS
*/
