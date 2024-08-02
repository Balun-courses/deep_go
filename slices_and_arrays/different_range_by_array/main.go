package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3}
	for value := range data { // copy of array
		fmt.Println(value)
	}

	for value := range &data { // not a copy of array
		fmt.Println(value)
	}

	for value := range data[:] { // not a copy of array
		fmt.Println(value)
	}
}
