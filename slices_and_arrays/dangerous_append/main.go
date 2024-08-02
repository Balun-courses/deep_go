package main

import "fmt"

func main() {
	data := make([]int, 4)
	copy := append(data, 5)

	fmt.Println(data)
	fmt.Println(copy)
}
