package main

import "fmt"

func divide(lhs, rhs int) {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	_ = lhs / rhs
}

func main() {
	divide(1000, 0)
}
