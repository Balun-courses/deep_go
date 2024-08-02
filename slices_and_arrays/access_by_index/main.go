package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data1 := [5]int8{1, 2, 3, 4, 5}
	pointer1 := uintptr(unsafe.Pointer(&data1))

	first1 := *(*int8)(unsafe.Pointer(pointer1))      // => data1[0]
	second1 := *(*int8)(unsafe.Pointer(pointer1 + 1)) // => data1[1]
	third1 := *(*int8)(unsafe.Pointer(pointer1 + 2))  // => data1[2]

	fmt.Println(first1, second1, third1)

	data2 := [5]int32{1, 2, 3, 4, 5}
	pointer2 := uintptr(unsafe.Pointer(&data2))

	first2 := *(*int32)(unsafe.Pointer(pointer2))                              // => data2[0]
	second2 := *(*int32)(unsafe.Pointer(pointer2 + 1*unsafe.Sizeof(int32(0)))) // => data2[1]
	third2 := *(*int32)(unsafe.Pointer(pointer2 + 2*unsafe.Sizeof(int32(0))))  // => data2[2]

	fmt.Println(first2, second2, third2)
}
