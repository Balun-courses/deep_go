package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}

	fmt.Println(data)
	process1(data)
	fmt.Println(data)
	process2(data)
	fmt.Println(data)
}

func process1(data []int) {
	data[2] = 5
}

func process2(data []int) {
	data = append(data, 6)
}
