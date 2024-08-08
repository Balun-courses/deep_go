package main

import (
	"testing"
)

// go test -bench=. clear_test.go

func BenchmarkClearWithFive(b *testing.B) {
	data := make([]int, 10*1024)
	for i := 0; i < b.N; i++ {
		for idx := range data {
			data[idx] = 5
		}
	}
}

func BenchmarkClearWithZero(b *testing.B) {
	data := make([]int, 10*1024)
	for i := 0; i < b.N; i++ {
		for idx := range data {
			data[idx] = 0
		}
	}
}
