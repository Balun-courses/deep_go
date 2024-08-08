package main

import "fmt"

func main() {
	var array *[4]int // = nil

	// _ = array[0] -> panic
	// array[0] = 1 -> panic

	fmt.Println("length =", len(array))
	fmt.Println("capacity =", cap(array))

	for idx := range array { // ok
		fmt.Println(idx)
	}

	for idx, value := range array { // panic
		fmt.Println(idx, value)
	}
}
