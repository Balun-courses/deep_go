package main

import (
	"fmt"
	"unsafe"
)

//go:noinline
func allocation(index int) byte {
	var data [1 << 20]byte
	return data[index]
}

func main() {
	var array [10]int
	address := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#1 array address:", address)

	allocation(100)

	address = (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#2 array address:", address)
}
