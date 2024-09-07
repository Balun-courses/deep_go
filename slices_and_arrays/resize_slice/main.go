package main

import "fmt"

func main() {
	data := make([]int, 3, 6) // len = 3, cap = 6

	// fmt.Println(data[4]) // panic

	data = data[:6] // len = 6, cap = 6

	fmt.Println(data[4])
}
