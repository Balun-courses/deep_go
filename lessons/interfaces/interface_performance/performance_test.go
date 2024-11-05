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

type Value2 struct {
	number int
}

//go:noinline
func (v Value2) Add() int {
	return v.number + v.number
}

var Result int

func Get(index int) ValueInterface {
	if index == 0 {
		return Value{}
	} else {
		return Value2{}
	}
}

func BenchmarkDirect(b *testing.B) {
	var value Value
	for i := 0; i < b.N; i++ {
		Result = value.Add()
	}
}

func BenchmarkWithInterface1(b *testing.B) {
	var iface ValueInterface = Get(0)
	for i := 0; i < b.N; i++ {
		Result = iface.Add()
	}
}

func BenchmarkWithInterface2(b *testing.B) {
	var iface ValueInterface = Get(1)
	for i := 0; i < b.N; i++ {
		Result = iface.(Value2).Add()
	}
}
