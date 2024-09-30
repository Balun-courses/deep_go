package main

import "fmt"

type Data1 struct {
	Value string
}

type Data2 struct {
	Value string
}

func Print1[T any](value T) {
	fmt.Println(value.Value) // compilation error
}

func Print2[T Data1 | Data2](value T) {
	fmt.Println(value.Value) // compilation error
}
