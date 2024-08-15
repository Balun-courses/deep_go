package main

import "testing"

// go test -bench=. comparison_test.go

//go:noinline
func withDefer(a, b int) {
	defer func() {
		Result = a + b
	}()
}

//go:noinline
func withoutDefer(a, b int) {
	Result = a + b
}

var Result int

func BenchmarkWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDefer(-1, i)
	}
}

func BenchmarkWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDefer(-1, i)
	}
}
