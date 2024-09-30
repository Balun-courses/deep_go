package main

import (
	"fmt"
	"unsafe"
)

type data struct{}

func main() {
	x := &data{}
	y := data{}

	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))
}
