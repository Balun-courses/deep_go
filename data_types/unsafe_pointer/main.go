package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value uint32 = 0xFFFFFFFF

	pointer := unsafe.Pointer(&value)
	bytePointer := (*uint8)(pointer)

	fmt.Println(*bytePointer)

	pointer = unsafe.Add(pointer, 2)
	twoBytePointer := (*uint16)(pointer)

	fmt.Println(*twoBytePointer)
}
