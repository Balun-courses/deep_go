package main

import "fmt"

func Decorator[T any](fn func(T), decorator func(T) T) func(T) {
	return func(input T) {
		fn(decorator(input))
	}
}

func main() {
	print := func(value int) { fmt.Println(value) }
	double := func(value int) int { return value * value }

	decorated := Decorator(print, double)
	decorated(10)
}
