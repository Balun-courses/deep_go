package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const elemSize = unsafe.Sizeof(int32(0)) // 4 bytes
	array := [...]int32{1, 2, 3}
	pointer := unsafe.Pointer(&array)

	first := *(*int32)(unsafe.Add(pointer, 0*elemSize))  // => data2[0]
	second := *(*int32)(unsafe.Add(pointer, 1*elemSize)) // => data2[1]
	third := *(*int32)(unsafe.Add(pointer, 2*elemSize))  // => data2[2]

	dangerous1 := *(*int32)(unsafe.Add(pointer, -1)) // => data2[-1]
	dangerous2 := *(*int32)(unsafe.Add(pointer, 3))  // => data2[3]

	fmt.Println("correct:", first, second, third)
	fmt.Println("dangeroues:", dangerous1, dangerous2)
}
