package main

import (
	"errors"
	"testing"

	othererrors "github.com/pkg/errors"
)

// go test -bench=. performance_test.go

var err error

func BenchmarkErrorWithoutStackTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = errors.New("error")
	}
}

func BenchmarkErrorWithStackTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = othererrors.New("error")
	}
}
