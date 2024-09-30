package main

import "fmt"

const (
	OkStatus = iota
	EvenNumberErr
	ZeroNumberErr
)

func divideV1(lhs, rhs int) (int, int) {
	if rhs == 0 {
		return 0, ZeroNumberErr
	} else if lhs%2 == 0 || rhs%2 == 0 {
		return 0, EvenNumberErr
	}

	return lhs / rhs, OkStatus
}

func divideV2(lhs, rhs int, status *int) int {
	if rhs == 0 {
		*status = ZeroNumberErr
		return 0
	} else if lhs%2 == 0 || rhs%2 == 0 {
		*status = EvenNumberErr
		return 0
	}

	*status = OkStatus
	return lhs / rhs
}

func main() {
	x := 100
	y := 0

	value, status := divideV1(x, y)
	fmt.Println(value, status)

	value = divideV2(x, y, &status)
	fmt.Println(value, status)
}
