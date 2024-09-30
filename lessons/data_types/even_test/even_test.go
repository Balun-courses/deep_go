package main

import (
	"testing"
)

// go test -bench=. even_test.go

var result int

func BenchmarkEven1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result = i & 1
		_ = result
	}
}

func BenchmarkEven2(b *testing.B) {
	for i := 1; i < b.N; i++ {
		result = i % 2
		_ = result
	}
}
