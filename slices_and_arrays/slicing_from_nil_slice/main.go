package main

import "fmt"

func main() {
	var data []int

	fmt.Println(data == nil)

	slice := data[:]

	fmt.Println(slice == nil)

	slice = data[0:0]

	fmt.Println(slice == nil)

	// slice = data[0:1]  ->  panic
}
