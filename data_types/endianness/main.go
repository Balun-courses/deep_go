package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var number int32 = 0x12345678
	pointer := uintptr(unsafe.Pointer(&number))

	fmt.Printf("0x")
	for i := 0; i < 4; i++ {
		offset := uintptr(i)
		byteValue := *(*int8)(unsafe.Pointer(pointer + offset))
		fmt.Printf("%x", byteValue)
	}

	fmt.Println()
}
