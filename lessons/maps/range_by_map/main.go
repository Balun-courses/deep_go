package main

import "fmt"

func main() {
	data := map[int]int{1: 1, 2: 2, 3: 3}
	for _, value := range data {
		value = 1000
		_ = value
	}

	fmt.Println(data)
}
