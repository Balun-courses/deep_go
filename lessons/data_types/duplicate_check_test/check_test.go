package main

import (
	"testing"
)

// go test -bench=. check_test.go

func HasDuplicatesFrom1To7WithBits(data []int) int {
	var lookup int8
	for _, number := range data {
		if lookup&(1<<number) == 1 {
			return number
		}

		lookup |= 1 << number
	}

	return -1
}

func HasDuplicatesFrom1To7WithArray(data []int) int {
	var lookup [8]int8
	for _, number := range data {
		if lookup[number] == 1 {
			return number
		}

		lookup[number] = 1
	}

	return -1
}

func HasDuplicatesFrom1To7WithHashTable(data []int) int {
	lookup := make(map[int]struct{}, len(data))
	for _, number := range data {
		_, found := lookup[number]
		if found {
			return number
		}

		lookup[number] = struct{}{}
	}

	return -1
}

func BenchmarkHasDuplicatesFrom1To7WithBits(b *testing.B) {
	data := []int{1, 5, 2, 3, 6, 4, 7, 2, 7}
	for i := 1; i < b.N; i++ {
		_ = HasDuplicatesFrom1To7WithBits(data)
	}
}

func BenchmarkHasDuplicatesFrom1To7WithArray(b *testing.B) {
	data := []int{1, 5, 2, 3, 6, 4, 7, 2, 7}
	for i := 1; i < b.N; i++ {
		_ = HasDuplicatesFrom1To7WithArray(data)
	}
}

func BenchmarkHasDuplicatesFrom1To7WithHashTable(b *testing.B) {
	data := []int{1, 5, 2, 3, 6, 4, 7, 2, 7}
	for i := 1; i < b.N; i++ {
		_ = HasDuplicatesFrom1To7WithHashTable(data)
	}
}
