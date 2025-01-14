package main

import "fmt"

func divide(lhs, rhs float32) float32 {
	return lhs / rhs
}

func main() {
	result := divide(1000, 0)
	fmt.Println(result)
}
