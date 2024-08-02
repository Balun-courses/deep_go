package main

import "fmt"

func main() {
	data := map[int]struct{}{
		1: {}, 2: {}, 3: {}, 4: {}, 5: {},
	}

	for key := range data {
		fmt.Print(key, " ")
	}

	fmt.Println()
	fmt.Println(data)
}
