package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x, y int = 1, 2
	)

	p1 := unsafe.Pointer(&x)
	p2 := unsafe.Add(p1, 8)

	fmt.Println("p1: ", *(*int)(p1))
	fmt.Println("p2: ", *(*int)(p2))

	_ = y
}
