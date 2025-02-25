package main

import (
	"testing"
	"unsafe"
)

// go test -gcflags="-d=ssa/check_bce" -bench=. bench_test.go

func sum1(values []int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += values[i]
	}

	return sum
}

func sum2(values []int, size int) int {
	sum := 0
	startPtr := unsafe.Pointer(&values[0])
	for i := 0; i < size-1; i++ {
		sum += *(*int)(
			unsafe.Pointer(uintptr(startPtr) + uintptr(i)*unsafe.Sizeof(values[0])),
		)
	}

	return sum
}

var Result int

func BenchmarkSum1(b *testing.B) {
	values := make([]int, 1<<20)
	for i := 0; i < b.N; i++ {
		Result = sum1(values, 1<<20)
	}
}

func BenchmarkSum2(b *testing.B) {
	values := make([]int, 1<<20)
	for i := 0; i < b.N; i++ {
		Result = sum2(values, 1<<20)
	}
}
