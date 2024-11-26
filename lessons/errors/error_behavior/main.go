package main

import (
	"fmt"
)

type ErrorFromAnotherPackage struct{}

func (e ErrorFromAnotherPackage) Error() string {
	return "error"
}

func (e ErrorFromAnotherPackage) Path() string {
	return "path"
}

type FSError interface {
	Path() string
}

func IsFSError(err error) bool {
	_, ok := err.(FSError)
	return ok
}

func getErrorFromAnotherPackage() error {
	return ErrorFromAnotherPackage{}
	//return fmt.Errorf("%w", ErrorFromAnotherPackage{})
}

func main() {
	err := getErrorFromAnotherPackage()
	if IsFSError(err) {
		fmt.Println("FSError")
	}
}
