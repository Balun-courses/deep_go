package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []byte("Hello world")
	strData := unsafe.String(unsafe.SliceData(data), len(data))

	fmt.Println(strData)
	data[0] = 'W'
	fmt.Println(strData)
}
