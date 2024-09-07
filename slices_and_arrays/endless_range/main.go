package main

import "fmt"

func main() {
	data := []int{0, 1, 2}
	for range data {
		data = append(data, 10)
		fmt.Println("iteration")
	}

	/*
		data := []int{0, 1, 2}
		for i := 0; i < len(data); i++ {
			data = append(data, 10)
			fmt.Println("iteration")
		}
	*/
}
