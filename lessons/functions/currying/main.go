package main

import "fmt"

func multiply(x int) func(y int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	fmt.Println(multiply(10)(15))

	// частичное применение
	var mult10 = multiply(10)
	var mult15 = multiply(15)

	fmt.Println(mult10(5))
	fmt.Println(mult15(15))
}
