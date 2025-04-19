package main

import "testing"

// go test -gcflags="-d=ssa/check_bce" -bench=. bench_test.go

func process1(data string) []byte {
	return []byte{
		data[0], // Found IsInBounds
		data[1], // Found IsInBounds
		data[2], // Found IsInBounds
		data[3], // Found IsInBounds
	}
}

func process2(data string) []byte {
	return []byte{
		data[3], // Found IsInBounds
		data[2],
		data[1],
		data[0],
	}
}

var Result []byte

func BenchmarkProcess1(b *testing.B) {
	data := "Hello world"
	for i := 0; i < b.N; i++ {
		Result = process1(data)
	}
}

func BenchmarkProcess2(b *testing.B) {
	data := "Hello world"
	for i := 0; i < b.N; i++ {
		Result = process2(data)
	}
}
