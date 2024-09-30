package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("exited #1")
	}()

	fmt.Println("started")

	defer func() {
		recover()
		fmt.Println("recovered")
	}()

	panic("exit")

	fmt.Println("exited #2")
}
