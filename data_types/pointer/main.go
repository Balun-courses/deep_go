package main

import "fmt"

func main() {
	var value int32 = 100
	var pointer *int32 = &value

	fmt.Println(pointer)
	fmt.Println(*pointer)

	*pointer = 500

	fmt.Println(pointer)
	fmt.Println(*pointer)
}
