package main

import "fmt"

func Create[T any]() *T {
	return new(T)
}

func Print[T any](value T) {
	fmt.Println()
}

func MultiplePrint[T1 any, T2 any](lhs T1, rhs T2) {
	fmt.Println(lhs, rhs)
}

type Data[T any] struct {
	Value T
}

func main() {
	// without infence
	value1 := Create()
	_ = value1
	var value2 *int = Create()
	_ = value2

	// with infence
	Print(100)
	Print("hello")
	MultiplePrint(100, "hello")

	// without infence
	data1 := Data{}
	_ = data1
	data2 := Data{Value: 100}
	_ = data2
}
