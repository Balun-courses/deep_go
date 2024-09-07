package main

import (
	"fmt"
	"unsafe"
)

func main() {
	values := [...]int{100, 200, 300}

	// also actual for slices
	for idx, value := range values {
		value += 50
		fmt.Println("#1:", unsafe.Pointer(&value), "#2:", unsafe.Pointer(&values[idx]))
	}

	fmt.Println("values:", values)
}
