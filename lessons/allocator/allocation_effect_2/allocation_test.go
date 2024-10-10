package main

import "testing"

// go test -bench=. -benchmem

const bufferSize = 10

type ReaderWithSliceArgument struct{}

func (r ReaderWithSliceArgument) Read(p []byte) (int, error) {
	for i := 0; i < bufferSize; i++ {
		p[i] = byte(i)
	}

	return bufferSize, nil
}

type ReaderWithSliceReturn struct{}

func (r ReaderWithSliceReturn) Read(n int) ([]byte, error) {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = byte(i)
	}

	return p, nil
}

func BenchmarkSliceWithArgument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := make([]byte, bufferSize)
		reader := ReaderWithSliceArgument{}
		_, _ = reader.Read(p)
	}
}

func BenchmarkSliceWithReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := ReaderWithSliceReturn{}
		_, _ = reader.Read(bufferSize)
	}
}
