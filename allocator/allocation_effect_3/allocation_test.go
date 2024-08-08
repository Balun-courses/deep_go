package main

import "testing"

// go test -bench=. -benchmem

type Data struct {
	pointer *int
}

func BenchmarkIteration1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var number int
		data := &Data{
			pointer: &number,
		}
		_ = data
	}
}

func BenchmarkIteration2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var number int
		data := &Data{}
		data.pointer = &number
		_ = data
	}
}
