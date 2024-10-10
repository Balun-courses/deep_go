package main

// go test -bench=. allocations_test.go -benchmem

import (
	"testing"
)

//go:noinline
func Initialize(value *int) {
	*value = 1000
}

func BenchmarkWithoutLoopAllocation(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		Initialize(&value)
	}
}

func BenchmarkWithLoopAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var value int
		Initialize(&value)
	}
}
