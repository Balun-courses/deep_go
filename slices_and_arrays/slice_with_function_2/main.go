package main

import "fmt"

func main() {
	data := make([]int, 3, 6)

	fmt.Println(data)
	process(data)
	fmt.Println(data)
}

func process(data []int) {
	data = append(data, 5)
}
