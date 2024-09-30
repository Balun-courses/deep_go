package main

import (
	"fmt"
)

func process() {
	fmt.Println("first")
	panic("error from process")
	fmt.Println("second")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()

	process()
}
