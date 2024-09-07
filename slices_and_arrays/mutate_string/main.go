package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []byte("Hello world")
	pointer := (*byte)(unsafe.SliceData(data))
	str := unsafe.String(pointer, len(data))

	fmt.Println("initial: ", str)
	data[0] = '#'
	fmt.Println("modified:", str)
}
