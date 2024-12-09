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
	d := [0]byte{}

	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&c))
	fmt.Println(unsafe.Pointer(&d))
}
