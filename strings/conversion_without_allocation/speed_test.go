package main

import (
	"testing"
	"unsafe"
)

// go test -bench=. speed_test.go

func BenchmarkConversion(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"
	for i := 0; i < b.N; i++ {
		data := []byte(text)
		_ = data
	}
}

func BenchmarkConversionOptimized(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"
	for i := 0; i < b.N; i++ {
		data := unsafe.Slice(unsafe.StringData(text), len(text))
		_ = data
	}
}
