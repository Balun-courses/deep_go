package main

import "fmt"

func main() {
	data := make([]int, 3, 6)

	// fmt.Println(data[4]) -> out of range

	data = data[:6]

	fmt.Println(data[4])
}
