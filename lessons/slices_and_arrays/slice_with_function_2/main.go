package main

import "fmt"

func main() {
	data := make([]int, 3, 6)

	fmt.Println("initial slice:", data)
	process(data)
	fmt.Println("after process:", data)
	fmt.Println("after process:", data[:4])
}

func process(data []int) {
	data = append(data, 5)
}
