package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []byte("Hello world")
	strData := *(*string)(unsafe.Pointer(&data))

	fmt.Println(strData)
	data[0] = 'h'
	fmt.Println(strData)

}
