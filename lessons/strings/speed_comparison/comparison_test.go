package main

import "testing"

// go test -bench=. comparison_test.go

func BenchmarkComparison(b *testing.B) {
	bs := make([]byte, 1<<26)
	s0 := string(bs)
	s1 := string(bs)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s0 == s1
	}
}

func BenchmarkComparisonOptimized(b *testing.B) {
	bs := make([]byte, 1<<26)
	s0 := string(bs)
	s1 := s0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s0 == s1
	}
}
