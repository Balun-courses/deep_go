package main

import "fmt"

// Need to show solution with maps.Clone

func main() {
	data := map[int]struct{}{
		0: {},
		1: {},
		2: {},
	}

	for key := range data {
		data[10+key] = struct{}{}
	}

	fmt.Println(data)
}
