package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

var (
	ErrNumber1 = errors.New("error 1")
	ErrNumber2 = errors.New("error 1")
	ErrNumber3 = errors.New("error 3")
)

// fully compatible with the Go standard library errors package
//     works for errors.Is, errors.As and errors.Unwrap

func main() {
	var err error
	err = multierror.Append(err, ErrNumber1)
	err = multierror.Append(err, ErrNumber2)
	err = fmt.Errorf("internal error: %w", err)

	if errors.Is(err, ErrNumber1) {
		fmt.Println("found error1 with errors.Is")
	}
	if errors.Is(err, ErrNumber2) {
		fmt.Println("found error2 with errors.Is")
	}
	if errors.Is(err, ErrNumber3) {
		fmt.Println("found error3 with errors.Is")
	}
}
