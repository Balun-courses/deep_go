package main

import (
	"errors"
	"fmt"
)

func wrap(err error) error {
	return fmt.Errorf("wrapped error: %w", err)
}

func main() {
	err := errors.New("internal error")
	fmt.Println(wrap(err).Error())
}
