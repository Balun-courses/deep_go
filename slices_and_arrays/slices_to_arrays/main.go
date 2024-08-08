package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])

	slice[0] = 10

	fmt.Println(array, len(array), cap(array), unsafe.Sizeof(array))

	// _ = [4]int(slice[:4]) -> ok
	// _ = [4]int(slice[:3]) -> panic
}
