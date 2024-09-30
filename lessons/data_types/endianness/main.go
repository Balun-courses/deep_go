package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var number int32 = 0x12345678
	pointer := unsafe.Pointer(&number)

	fmt.Printf("0x")
	for i := 0; i < 4; i++ {
		byteValue := *(*int8)(unsafe.Add(pointer, i))
		fmt.Printf("%x", byteValue)
	}

	fmt.Println()
}
