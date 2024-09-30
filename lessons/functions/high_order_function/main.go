package main

import "fmt"

func forEach(data []int, fn func(int) int) []int {
	newData := make([]int, 0, len(data))
	for _, value := range data {
		newData = append(newData, fn(value))
	}

	return newData
}

func main() {
	var data = []int{1, 2, 3, 4}
	var newData = forEach(data, func(value int) int {
		return (value * value)
	})

	fmt.Println(data)
	fmt.Println(newData)
}
