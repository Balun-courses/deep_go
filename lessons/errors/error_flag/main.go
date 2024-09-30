package main

import "fmt"

func divideV1(lhs, rhs int) (int, bool) {
	if rhs == 0 {
		return 0, false
	}

	return lhs / rhs, true
}

func divideV2(lhs, rhs int, status *bool) int {
	*status = false
	if rhs == 0 {
		return 0
	}

	*status = true
	return lhs / rhs
}

func main() {
	x := 100
	y := 0

	value, ok := divideV1(x, y)
	fmt.Println(value, ok)

	value = divideV2(x, y, &ok)
	fmt.Println(value, ok)
}
