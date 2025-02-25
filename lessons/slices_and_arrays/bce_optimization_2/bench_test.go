package main

import "testing"

// go test -gcflags="-d=ssa/check_bce" -bench=. bench_test.go

func fn1(s []int, i int) {
	_ = s[i+3] // Found IsInBounds
	_ = s[i+2] // Found IsInBounds
	_ = s[i+1] // Found IsInBounds
	_ = s[i]   // Found IsInBounds
}

func fn2(s []int, i int) {
	s = s[i : i+4] // Found IsSliceInBounds
	_ = s[3]
	_ = s[2]
	_ = s[1]
	_ = s[0]
}

func BenchmarkFN1(b *testing.B) {
	values := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		fn1(values, 0)
	}
}

func BenchmarkFN2(b *testing.B) {
	values := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		fn2(values, 0)
	}
}
