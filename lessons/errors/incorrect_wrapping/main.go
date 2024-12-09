package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("source error 1")
	err2 := errors.New("source error 2")
	err := fmt.Errorf("additional error information: %w and %w", err1, err2)

	fmt.Println(err.Error())
	fmt.Println(errors.Unwrap(err))

	err = fmt.Errorf("additional error information: %w", "error")

	fmt.Println(err.Error())
	fmt.Println(errors.Unwrap(err))
}
