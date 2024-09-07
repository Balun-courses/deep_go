package main

import "fmt"

func main() {
	var array *[4]int // = nil

	fmt.Println("length =", len(array))
	fmt.Println("capacity =", cap(array))

	for idx := range array { // ok
		fmt.Println(idx)
	}

	for idx, value := range array { // panic
		fmt.Println(idx, value)
	}
}
