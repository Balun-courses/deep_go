package main

import "fmt"

func process1() {
	defer fmt.Println("test1")
	defer fmt.Println("test2")

	// implementation...
}

func process2() {
	defer func() {
		fmt.Println("test2")
		fmt.Println("test1")
	}()

	// implementation...
}
