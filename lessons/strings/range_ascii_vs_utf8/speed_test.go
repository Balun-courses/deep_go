package main

import (
	"testing"
)

// go test -bench=. speed_test.go

func BenchmarkByRangeBytes(b *testing.B) {
	text := []byte("aaaabbbbccccddddeeeeffff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range text {
		}
	}
}

func BenchmarkByRangeStringUTF8(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range text {
		}
	}
}
