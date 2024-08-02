package main

import (
	"errors"
	"fmt"
)

func wrap(err error) error {
	return fmt.Errorf("failed to wrap: %w", err)
}

func main() {
	err := errors.New("error")
	fmt.Println(wrap(err).Error())
}
