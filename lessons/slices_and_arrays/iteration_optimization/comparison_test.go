package main

import (
	"testing"
)

// go test -bench=. comparison_test.go

const size = 1 << 10

func BenchmarkWithoutOptimizations(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data); j++ {
			data[j] = i
		}
	}
}

func BenchmarkWithCalculatedLength(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		length := len(data)
		for j := 0; j < length; j++ {
			data[j] = i
		}
	}
}

func BenchmarkWithLoopUnwinding(b *testing.B) {
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data)/4; j += 4 {
			data[j] = i
			data[j+1] = i
			data[j+2] = i
			data[j+3] = i
		}
	}
}
