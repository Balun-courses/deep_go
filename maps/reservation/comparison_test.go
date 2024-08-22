package main

import "testing"

// go test -bench=. comparison_test.go

func BenchmarkWithoutReservation(b *testing.B) {
	sourseData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 0; i < b.N; i++ {
		lookup := make(map[int]struct{})
		for j := 0; j < len(sourseData); j++ {
			lookup[j] = struct{}{}
		}
	}
}

func BenchmarkWithReservation(b *testing.B) {
	sourseData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for i := 0; i < b.N; i++ {
		lookup := make(map[int]struct{}, len(sourseData))
		for j := 0; j < len(sourseData); j++ {
			lookup[j] = struct{}{}
		}
	}
}
