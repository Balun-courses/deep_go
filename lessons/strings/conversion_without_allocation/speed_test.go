package main

import (
	"testing"
	"unsafe"
)

// go test -bench=. speed_test.go

var Result []byte

func BenchmarkConversion(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"
	for i := 0; i < b.N; i++ {
		Result = []byte(text)
	}
}

func BenchmarkConversionOptimized(b *testing.B) {
	text := "aaaabbbbccccddddeeeeffff"
	for i := 0; i < b.N; i++ {
		Result = unsafe.Slice(unsafe.StringData(text), len(text))
	}
}
