package main

import (
	"fmt"
	"unsafe"
)

type empty struct{}

func main() {
	a := struct{}{}
	b := struct{}{}
	c := empty{}

	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&c))
}
