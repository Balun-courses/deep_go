package main

import "fmt"

func process(temp *int32) {
	var value2 int32 = 200
	temp = &value2
}

func main() {
	var value1 int32 = 100
	pointer := &value1

	fmt.Println(*pointer)
	fmt.Println(pointer)

	process(pointer)

	fmt.Println(*pointer)
	fmt.Println(pointer)
}
