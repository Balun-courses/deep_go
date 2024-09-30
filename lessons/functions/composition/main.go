package main

import "fmt"

func sqr(number int) int {
	return number * number
}

func neg(number int) int {
	return -number
}

func compose(fn ...func(int) int) func(int) int {
	return func(value int) int {
		for _, v := range fn {
			value = v(value)
		}

		return value
	}
}

func main() {
	fn := compose(sqr, neg, sqr)
	fmt.Println(fn(4))
}
