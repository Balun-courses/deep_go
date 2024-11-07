package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("source error")
	err = fmt.Errorf("additional error information: %v", err)

	fmt.Println(err.Error())
	fmt.Println(errors.Unwrap(err))
}
