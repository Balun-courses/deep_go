package main

import "fmt"

func recursion() {
	var data = [10 * 1024 * 1024]int8{}
	_ = data

	recursion()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	recursion()
}
