package main

import "fmt"

func process(temp **int32) {
	var value2 int32 = 200
	*temp = &value2
}

func main() {
	var value1 int32 = 100
	pointer := &value1

	fmt.Println("value:", *pointer)
	fmt.Println("address:", pointer)

	process(&pointer)

	fmt.Println("value", *pointer)
	fmt.Println("address:", pointer)
}
