package main

import (
	"strconv"
	"testing"
)

var (
	n = 100_000
)

func BenchmarkMapBool(b *testing.B) {
	m := make(map[string]bool, n)

	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			m[strconv.Itoa(i)] = true
		}

		for i := 0; i < n; i++ {
			_, ok := m[strconv.Itoa(i)]
			_ = ok
		}
	}
}

func BenchmarkMapStruct(b *testing.B) {
	m := make(map[string]struct{}, n)

	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			m[strconv.Itoa(i)] = struct{}{}
		}

		for i := 0; i < n; i++ {
			_, ok := m[strconv.Itoa(i)]
			_ = ok
		}
	}
}
