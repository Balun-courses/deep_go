package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Mul(x, y int) int {
	return x * y
}

func Calculate(x, y int, fn func(int, int) int) int {
	fmt.Printf("x=%d y=%d\n", x, y)
	return fn(x, y)
}

func CalculateAdd(x, y int) int {
	fmt.Printf("x=%d y=%d\n", x, y)
	return Add(x, y)
}

func main() {
	Calculate(10, 10, Add)
	Calculate(10, 10, Mul)

	CalculateAdd(10, 10)
}
