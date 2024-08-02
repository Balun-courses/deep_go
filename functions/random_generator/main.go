package main

import (
	"fmt"
)

func produce(source int, permutation func(int) int) func() int {
	return func() int {
		source = permutation(source)
		return source
	}
}

func mutate(number int) int {
	return (1664525*number + 1013904223) % 2147483647
}

func main() {
	next := produce(1, mutate)

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
