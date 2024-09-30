package main

import "fmt"

func main() {
	// remove all elements
	first := map[int]int{1: 1, 2: 2, 3: 3}
	first = nil
	fmt.Println(first, " : ", len(first))

	// keep allocated memory
	second := map[int]int{1: 1, 2: 2, 3: 3}
	for key := range second {
		delete(second, key)
	}
	fmt.Println(second, " : ", len(second))

	// keep allocated memory
	third := map[int]int{1: 1, 2: 2, 3: 3}
	clear(third)
	fmt.Println(third, " : ", len(third))
}
