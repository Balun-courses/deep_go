package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "error"
}

func main() {
	var pointer *MyError = nil
	var err error = pointer
	fmt.Println("nil:", err == nil)
}
