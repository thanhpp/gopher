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
	rand.Seed(102034)

	sliceA := make([]A, n)

	for i := 0; i < n; i++ {
		sliceA = append(sliceA, A{
			I: rand.Intn(n),
		})
	}

	return sliceA
}

func BenchmarkNonGenericStableSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := genData()
		b.StartTimer()
		sort.SliceStable(data, func(i, j int) bool {
			return data[i].I < data[j].I
		})
	}
}

func BenchmarkNonGenericSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := genData()
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i].I < data[j].I
		})
	}
}

func BenchmarkGenericStableSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := genData()
		b.StartTimer()
		slices.SortStableFunc(data, func(a, b A) bool {
			return a.I < b.I
		})
	}
}

func BenchmarkGenericSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := genData()
		b.StartTimer()
		slices.SortFunc(data, func(a, b A) bool {
			return a.I < b.I
		})
	}
}

/*
- add stop/start timer

goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-generic-sort
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNonGenericStableSort
BenchmarkNonGenericStableSort-12    	     520	   2907970 ns/op	      56 B/op	       2 allocs/op
BenchmarkNonGenericSort
BenchmarkNonGenericSort-12          	    1161	   1596522 ns/op	      56 B/op	       2 allocs/op
BenchmarkGenericStableSort
BenchmarkGenericStableSort-12       	     625	   1944742 ns/op	       0 B/op	       0 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12             	    1442	   1406962 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-generic-sort	8.268s
*/

/*
- same rand seed

goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-generic-sort
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNonGenericStableSort
BenchmarkNonGenericStableSort-12    	     555	   2299454 ns/op	  507963 B/op	       6 allocs/op
BenchmarkNonGenericSort
BenchmarkNonGenericSort-12          	    1020	   1262832 ns/op	  507963 B/op	       6 allocs/op
BenchmarkGenericStableSort
BenchmarkGenericStableSort-12       	     729	   1895796 ns/op	  507907 B/op	       4 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12             	    1042	   1050101 ns/op	  507907 B/op	       4 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-generic-sort	5.664s
*/

/*
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/cmd/bench-generic-sort
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkNonGenericStableSort
BenchmarkNonGenericStableSort-12    	     565	   2561965 ns/op	  507961 B/op	       6 allocs/op
BenchmarkNonGenericSort
BenchmarkNonGenericSort-12          	     819	   1429645 ns/op	  507961 B/op	       6 allocs/op
BenchmarkGenericStableSort
BenchmarkGenericStableSort-12       	     532	   2109854 ns/op	  507905 B/op	       4 allocs/op
BenchmarkGenericSort
BenchmarkGenericSort-12             	     930	   1231374 ns/op	  507905 B/op	       4 allocs/op
PASS
ok  	github.com/thanhpp/gopher/cmd/bench-generic-sort	5.614s
*/

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
