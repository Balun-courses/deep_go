package main

import (
	"strings"
	"testing"
)

// go test -bench=. speed_test.go

func BenchmarkSimpleConcatenation(b *testing.B) {
	str := "test"
	for i := 0; i < b.N; i++ {
		str += "test"
	}

	_ = str
}

func BenchmarkConcatenationWithStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}

func BenchmarkConcatenationWithStringBuilderOptimized(b *testing.B) {
	builder := strings.Builder{}
	builder.Grow(4 + b.N*4)
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}
