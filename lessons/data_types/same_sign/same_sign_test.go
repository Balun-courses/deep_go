package main

import (
	"testing"
)

// go test -bench=. same_sign_test.go

func SlowSameSign(x, y int64) bool {
	return ((x < 0) && (y < 0)) || ((x >= 0) && (y >= 0))
}

func FastSameSign(x, y int64) bool {
	return (x ^ y) >= 0
}

var result bool

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = SlowSameSign(100, 200)
		_ = result
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = FastSameSign(100, 200)
		_ = result
	}
}
