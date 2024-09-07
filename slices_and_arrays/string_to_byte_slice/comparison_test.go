package main

import (
	"testing"
	"unsafe"
)

// go test -bench=. -benchmem comparison_test.go

func Convert(str string) []byte {
	if len(str) == 0 {
		return nil
	}

	return unsafe.Slice(unsafe.StringData(str), len(str))
}

var Result []byte

func BenchmarkConvertion(b *testing.B) {
	str := "Hello world!!!"
	for i := 0; i < b.N; i++ {
		Result = []byte(str)
	}
}

func BenchmarkUnsafeConvertion(b *testing.B) {
	str := "Hello world!!!"
	for i := 0; i < b.N; i++ {
		Result = Convert(str)
	}
}
