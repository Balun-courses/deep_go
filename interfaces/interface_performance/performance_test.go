package main

import "testing"

type ValueInterface interface {
	Add() int
}

type Value struct {
	number int
}

func (v Value) Add() int {
	return v.number + v.number
}

func BenchmarkDirect(b *testing.B) {
	value := Value{454}
	for i := 0; i < b.N; i++ {
		value.Add()
	}
}

func BenchmarkWithInterface(b *testing.B) {
	iface := ValueInterface(Value{454})
	for i := 0; i < b.N; i++ {
		iface.Add()
	}
}
