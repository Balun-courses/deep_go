package main

import "testing"

// go test -bench=. -benchmem performance_test.go

type ValueInterface interface {
	Add() int
}

type Value struct {
	number int
}

//go:noinline
func (v Value) Add() int {
	return v.number + v.number
}

var Result int

func BenchmarkDirect(b *testing.B) {
	var value Value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = value.Add()
	}
}

func BenchmarkWithInterface1(b *testing.B) {
	var value Value
	var iface ValueInterface = value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = iface.Add()
	}
}

func BenchmarkWithInterface2(b *testing.B) {
	var value Value
	var iface ValueInterface = value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = iface.(Value).Add()
	}
}
