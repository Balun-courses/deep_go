package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value1 int8 = 10
	fmt.Println("size1:", unsafe.Sizeof(value1)) // compile time

	var value2 int32 = 10
	fmt.Println("size2:", unsafe.Sizeof(value2)) // compile time
}
