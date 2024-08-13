package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "hello"
	pointer := unsafe.StringData(str[4:5])
	fmt.Println(pointer, string(*pointer))
}
