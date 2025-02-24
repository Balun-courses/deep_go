package main

import (
	"fmt"
	"unsafe"
)

// go run main.go
// go run -gcflags=-d=checkptr main.go

//go:noinline
func allocation(index int) byte {
	var data [1 << 20]byte
	return data[index]
}

func main() {
	var array [10]int
	address1 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("address1:", address1)

	allocation(100)

	address2 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("address2:", address2)
	fmt.Println("address1:", address1)
}
