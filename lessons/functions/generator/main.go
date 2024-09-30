package main

import "fmt"

func Generator(number int) func() int {
	return func() int {
		r := number
		number++
		return r
	}
}

func main() {
	generator := Generator(100)
	for i := 0; i <= 200; i++ {
		fmt.Println(generator())
	}
}
