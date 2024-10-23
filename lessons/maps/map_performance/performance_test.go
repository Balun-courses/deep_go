package main

import "testing"

// go test -bench=. performance_test.go

func BenchmarkWithSlice(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		value := slice[1]
		slice[1] = value
	}
}

func BenchmarkWithHashTable(b *testing.B) {
	table := map[int]int{0: 1, 1: 2, 2: 3}
	for i := 0; i < b.N; i++ {
		value := table[1]
		table[1] = value
	}
}
