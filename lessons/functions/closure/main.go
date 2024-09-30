package main

import "fmt"

func multiply(x int) func(y int) int {
	return func(y int) int {
		return x * y
	}
}

// Карринг функции — это изменение функции от вида func(a,b,c) до вида func(a)(b)(c).

func main() {
	// каррирование
	var mult10 = multiply(10)
	var mult15 = multiply(15)

	fmt.Println(mult10(5))
	fmt.Println(mult15(15))
}
