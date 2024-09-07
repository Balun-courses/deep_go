package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])

	slice[0] = 10

	fmt.Println("slice =", slice, len(slice), cap(slice))
	fmt.Println("array =", array, len(array), cap(array))
}
