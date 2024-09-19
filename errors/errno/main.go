package main

import "fmt"

const (
	OkStatus = iota
	EvenNumberErr
	ZeroNumberErr
)

var Errno = OkStatus

func divide(lhs, rhs int) int {
	if rhs == 0 {
		Errno = ZeroNumberErr
		return 0
	} else if lhs%2 == 0 || rhs%2 == 0 {
		Errno = EvenNumberErr
		return 0
	}

	Errno = OkStatus
	return lhs / rhs
}

func main() {
	x := 100
	y := 0

	value := divide(x, y)
	fmt.Println(value, Errno)
}
