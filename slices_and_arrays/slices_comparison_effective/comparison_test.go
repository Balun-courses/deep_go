package main

import (
	"fmt"
	"reflect"
	"testing"
)

// go test -bench=. -benchmem comparison_test.go

func equal(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func BenchmarkWithEqualFunction(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		_ = equal(lhs, rhs) // slices.Equal(...)
	}
}

func BenchmarkWithReflect(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(lhs, rhs)
	}
}

func BenchmarkWithSprint(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(lhs) == fmt.Sprint(rhs)
	}
}
