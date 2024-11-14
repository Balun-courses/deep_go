package main

import (
	"testing"
)

// go test -bench=. comparison_test.go

type Data1 struct {
	size   int32
	values [10 << 20]byte
}

type Data2 struct {
	values [10 << 20]byte
	size   int32
}

func BenchmarkComparisonData1(b *testing.B) {
	data1 := Data1{size: 100}
	data2 := Data1{size: 101}
	for i := 0; i < b.N; i++ {
		_ = data1 == data2
	}
}

func BenchmarkComparisonData2(b *testing.B) {
	data1 := Data2{size: 100}
	data2 := Data2{size: 101}
	for i := 0; i < b.N; i++ {
		_ = data1 == data2
	}
}
