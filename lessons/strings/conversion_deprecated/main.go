package main

import (
	"fmt"
	"unsafe"
)

/*
type SliceHeader struct {
    Data unsafe.Pointer
    Len  int
    Cap  int
}

type StringHeader struct {
    Data unsafe.Pointer
    Len  int
}
*/

func ToStringDeprecated(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func main() {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	str := ToStringDeprecated(data)
	fmt.Println(str)
}
