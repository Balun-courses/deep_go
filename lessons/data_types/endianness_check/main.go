package main

import (
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	var number int16 = 0x0001
	pointer := (*int8)(unsafe.Pointer(&number))
	return *pointer == 1
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}

func main() {
	if IsLittleEndian() {
		fmt.Println("Little endian")
	} else {
		fmt.Println("Big endian")
	}
}
