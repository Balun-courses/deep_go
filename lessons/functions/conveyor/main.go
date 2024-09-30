package main

import "fmt"

func sqr(number int) int {
	return number * number
}

func neg(number int) int {
	return -number
}

func inc(number int) int {
	return number + 1
}

func pipe(value int, fn ...func(int) int) int {
	for _, v := range fn {
		value = v(value)
	}

	return value
}

func reverse(fn ...func(int) int) []func(int) int {
	for idx := 0; idx < len(fn)/2; idx++ {
		fn[idx], fn[len(fn)-idx-1] = fn[len(fn)-idx-1], fn[idx]
	}

	return fn
}

func main() {
	// decorator way
	decorationResult := inc(neg(sqr(5)))
	fmt.Println(decorationResult)

	// composition way
	compositionResult1 := pipe(5, sqr, neg, inc)
	compositionResult2 := pipe(5, reverse(inc, neg, sqr)...)
	fmt.Println(compositionResult1, compositionResult2)
}
