package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("source error")
	err = fmt.Errorf("additional error information: %w", err)
	err = fmt.Errorf("internal error: %w", err)

	fmt.Println(err.Error())
	fmt.Println(errors.Unwrap(err))
	fmt.Println(errors.Unwrap(errors.Unwrap(err)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err))))
}
