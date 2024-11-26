package main

import (
	"fmt"
	"strconv"
)

type error interface {
	Error() string
}

type DivisionError struct {
	errMessage     string
	firstArgument  int
	secondArgument int
}

func NewDivisionError(message string, lhs, rhs int) DivisionError {
	return DivisionError{
		errMessage:     message,
		firstArgument:  lhs,
		secondArgument: rhs,
	}
}

func (e DivisionError) Error() string {
	if e.errMessage != "" {
		return e.errMessage + " [" + strconv.Itoa(e.firstArgument) + ", " + strconv.Itoa(e.secondArgument) + "]"
	} else {
		return ""
	}
}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, NewDivisionError("zero number", lhs, rhs)
	} else if lhs%2 == 0 || rhs%2 == 0 {
		return 0, NewDivisionError("even number", lhs, rhs)
	}

	return lhs / rhs, nil
}

func main() {
	x := 100
	y := 0

	value, err := divide(x, y)
	fmt.Println(value)
	fmt.Println(err)
}
