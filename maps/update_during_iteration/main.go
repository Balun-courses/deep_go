package main

import "fmt"

// Need to show solution

func main() {
	data := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for key, value := range data {
		if value {
			data[10+key] = value
		}
	}

	fmt.Println(data)
}
