package main

import "testing"

// go test -gcflags="-d=ssa/check_bce" -bench=. bench_test.go

func sum1(values []int) int {
	return values[0] + // Found IsInBounds
		values[1] + // Found IsInBounds
		values[2] + // Found IsInBounds
		values[3] // Found IsInBounds
}

func sum2(values []int) int {
	return values[3] + // Found IsInBounds
		values[2] +
		values[1] +
		values[0]
}

var Result int

func BenchmarkSum1(b *testing.B) {
	values := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		Result = sum1(values)
	}
}

func BenchmarkSum2(b *testing.B) {
	values := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		Result = sum2(values)
	}
}
