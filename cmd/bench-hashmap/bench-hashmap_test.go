package benchhashmap_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"golang.org/x/exp/constraints"
)

var (
	sizes = []int{1, 5, 10, 20, 30, 40, 50}
)

func BenchmarkSliceInt(b *testing.B) {
	for _, size := range sizes {
		// generate input
		s := make([]int, size)
		for i := 0; i < size; i++ {
			s[i] = i
		}
		b.ResetTimer()
		b.Run(fmt.Sprintf("size_%d", len(s)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					find := rand.Intn(size)
					benchSlice(s, find)
				}
			},
		)
	}
}

func BenchmarkMapInt(b *testing.B) {
	for _, size := range sizes {
		// generate input
		s := make(map[int]bool, size)
		for i := 0; i < size; i++ {
			s[i] = true
		}
		b.ResetTimer()
		b.Run(fmt.Sprintf("size_%d", len(s)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					find := rand.Intn(size)
					benchMap(s, find)
				}
			},
		)
	}
}

func BenchmarkSliceString(b *testing.B) {
	for _, size := range sizes {
		// generate input
		s := make([]string, size)
		for i := 0; i < size; i++ {
			s[i] = strconv.Itoa(i)
		}
		finds := make([]string, 0, 10_000)
		for i := 0; i < 10_000; i++ {
			finds = append(finds, strconv.Itoa(rand.Intn(size)))
		}

		b.ResetTimer()
		b.Run(fmt.Sprintf("size_%d", len(s)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					benchSlice(s, finds[i])
				}
			},
		)
	}
}

func BenchmarkMapString(b *testing.B) {
	for _, size := range sizes {
		// generate input
		s := make(map[string]bool, size)
		for i := 0; i < size; i++ {
			s[strconv.Itoa(i)] = true
		}
		finds := make([]string, 0, 10_000)
		for i := 0; i < 10_000; i++ {
			finds = append(finds, strconv.Itoa(rand.Intn(size)))
		}

		b.ResetTimer()
		b.Run(fmt.Sprintf("size_%d", len(s)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					benchMap(s, finds[i])
				}
			},
		)
	}
}

func benchSlice[T constraints.Ordered](input []T, find T) {
	for i := range input {
		if input[i] == find {
			break
		}
	}
}

func benchMap[T constraints.Ordered](input map[T]bool, find T) {
	_, _ = input[find]
}

/*
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/bench-hashmap
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkSliceInt/size_1-12             125404327                9.605 ns/op           0 B/op          0 allocs/op
BenchmarkSliceInt/size_5-12             69207064                17.54 ns/op            0 B/op          0 allocs/op
BenchmarkSliceInt/size_10-12            63447366                18.60 ns/op            0 B/op          0 allocs/op
BenchmarkSliceInt/size_20-12            58960952                20.43 ns/op            0 B/op          0 allocs/op
BenchmarkSliceInt/size_30-12            54392230                21.73 ns/op            0 B/op          0 allocs/op
BenchmarkSliceInt/size_40-12            53013524                23.26 ns/op            0 B/op          0 allocs/op
BenchmarkSliceInt/size_50-12            49305501                24.35 ns/op            0 B/op          0 allocs/op

BenchmarkMapInt/size_1-12               106518439               11.28 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_5-12               67979809                18.07 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_10-12              47070580                26.61 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_20-12              45591961                26.35 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_30-12              44955747                27.24 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_40-12              46008104                25.74 ns/op            0 B/op          0 allocs/op
BenchmarkMapInt/size_50-12              42991318                27.75 ns/op            0 B/op          0 allocs/op

BenchmarkSliceString/size_1-12          94805643                12.54 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_5-12          53173370                22.47 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_10-12         42057357                28.63 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_20-12         34653843                34.05 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_30-12         28817215                41.65 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_40-12         23304921                52.29 ns/op            0 B/op          0 allocs/op
BenchmarkSliceString/size_50-12         19114158                62.58 ns/op            0 B/op          0 allocs/op

BenchmarkMapString/size_1-12            93007335                13.08 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_5-12            52409577                23.35 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_10-12           42127460                26.81 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_20-12           39925528                30.12 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_30-12           43890310                27.75 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_40-12           39024145                31.18 ns/op            0 B/op          0 allocs/op
BenchmarkMapString/size_50-12           37039489                32.53 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/thanhpp/gopher/bench-hashmap 45.370s
*/

/*
go test -bench=. -benchtime=10000x
goos: linux
goarch: amd64
pkg: github.com/thanhpp/gopher/bench-hashmap
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkSliceInt/size_1-12                10000                10.16 ns/op
BenchmarkSliceInt/size_5-12                10000                17.88 ns/op
BenchmarkSliceInt/size_10-12               10000                18.74 ns/op
BenchmarkSliceInt/size_20-12               10000                20.00 ns/op
BenchmarkSliceInt/size_30-12               10000                23.17 ns/op
BenchmarkSliceInt/size_40-12               10000                23.75 ns/op
BenchmarkSliceInt/size_50-12               10000                24.68 ns/op

BenchmarkMapInt/size_1-12                  10000                11.36 ns/op
BenchmarkMapInt/size_5-12                  10000                18.35 ns/op
BenchmarkMapInt/size_10-12                 10000                26.76 ns/op
BenchmarkMapInt/size_20-12                 10000                27.34 ns/op
BenchmarkMapInt/size_30-12                 10000                27.49 ns/op
BenchmarkMapInt/size_40-12                 10000                25.81 ns/op
BenchmarkMapInt/size_50-12                 10000                26.76 ns/op

BenchmarkSliceString/size_1-12             10000                 3.457 ns/op
BenchmarkSliceString/size_5-12             10000                12.33 ns/op
BenchmarkSliceString/size_10-12            10000                25.57 ns/op
BenchmarkSliceString/size_20-12            10000                25.38 ns/op
BenchmarkSliceString/size_30-12            10000                35.37 ns/op
BenchmarkSliceString/size_40-12            10000                47.57 ns/op
BenchmarkSliceString/size_50-12            10000                62.01 ns/op

BenchmarkMapString/size_1-12               10000                 3.290 ns/op
BenchmarkMapString/size_5-12               10000                16.71 ns/op
BenchmarkMapString/size_10-12              10000                18.79 ns/op
BenchmarkMapString/size_20-12              10000                18.67 ns/op
BenchmarkMapString/size_30-12              10000                18.97 ns/op
BenchmarkMapString/size_40-12              10000                20.77 ns/op
BenchmarkMapString/size_50-12              10000                26.41 ns/op
PASS
ok      github.com/thanhpp/gopher/bench-hashmap 0.024s
*/
