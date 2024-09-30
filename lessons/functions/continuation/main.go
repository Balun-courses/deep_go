package main

import "fmt"

func Divide1(lhs, rhs int) int {
	if rhs == 0 {
		panic("incorrect argument") // Why???
	} else {
		return lhs / rhs
	}
}

func Divide2(lhs, rhs int, successFn func(int), errorFn func()) {
	if rhs == 0 {
		errorFn()
	} else {
		successFn(lhs / rhs)
	}
}

func main() {
	Divide2(100, 10,
		func(number int) {
			fmt.Println(number)
		},
		func() {
			fmt.Println("incorrect argument")
		},
	)
}
