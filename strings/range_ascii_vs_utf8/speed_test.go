package main

import (
	"testing"
)

func BenchmarkASCII(b *testing.B) {
	text := []byte("aaaabbbbccccddddeeeeffff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range text {
		}
	}
}

func BenchmarkUTF8(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range text {
		}
	}
}
