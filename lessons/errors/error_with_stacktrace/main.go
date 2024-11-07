package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	value, err := DoSomething()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Println(value)
}

func DoSomething() (string, error) {
	return "", errors.New("some error explanation here")
}
