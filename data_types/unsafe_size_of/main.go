package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value1 int8 = 10
	fmt.Println(unsafe.Sizeof(value1))

	var value2 int32 = 10
	fmt.Println(unsafe.Sizeof(value2))

}
