package main

import "golang.org/x/text/number"

// #1 scenario
type T[T1 any, T2 any] struct{}

type A = T[int, string]        // ok
type B[T1 any] = T[T1, string] // compilation error

// #2 scenario
type Derived[Base any] struct {
	Base // compilation error
}

// #3 scenario
type Data struct{}

func (d Data) Do[T any](T value) {} // compilation error

// #4 scenario
func Do[T ~int](number T) {
	var variable T = number   // ok
	const constant T = number // compilation error
}

// #5 scenario
type Data[SIZE int] struct { // size like a constant number
	arrray [SIZE]int // compilation error
}