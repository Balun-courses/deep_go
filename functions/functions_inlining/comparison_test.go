package main

import "testing"

// go test -bench=. comparison_test.go

//go:noinline
func maxWithoutInline(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWithInline(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMaxWithoutInline(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = maxWithoutInline(-1, i)
	}
	_ = r
}

func BenchmarkMaxWithInline(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = maxWithInline(-1, i)
	}
	_ = r
}
