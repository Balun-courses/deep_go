package main

import (
	"testing"
	"unsafe"
)

// go test -bench=. -benchmem performance_test.go

type Int int

var convertedData []int

func BenchmarkCast(b *testing.B) {
	data := make([]Int, 1024)
	for i := 0; i < b.N; i++ {
		convertedData = make([]int, 1024)
		for idx, value := range data {
			convertedData[idx] = int(value)
		}
	}
}

func BenchmarkUnsafeCast(b *testing.B) {
	data := make([]Int, 1024)
	for i := 0; i < b.N; i++ {
		convertedData = *(*[]int)(unsafe.Pointer(&data))
	}
}
