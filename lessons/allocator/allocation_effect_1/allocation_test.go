package main

import "testing"

// go test -bench=. -benchmem

type Data struct {
	iValue int
	sValue string
	bValue bool
}

//go:noinline
func NewDataByValue() Data {
	return Data{iValue: 100, sValue: "100", bValue: true}
}

//go:noinline
func NewDataByPointer() *Data {
	return &Data{iValue: 100, sValue: "100", bValue: true}
}

func BenchmarkNewByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDataByValue()
	}
}

func BenchmarkNewByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewDataByPointer()
	}
}
