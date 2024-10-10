package main

import (
	"fmt"
	"unsafe"
)

// For memory blocks larger than 32768 bytes,
// each of them is always composed of multiple
// memory pages. The memory page size used by the
// official standard Go runtime (1.22 versions) is 8192 bytes.

func main() {
	var data1 [32769]byte
	var data2 [32769]byte

	data1Pointer := unsafe.Pointer(&data1)
	data2Pointer := unsafe.Pointer(&data2)

	fmt.Println("adress:", data1Pointer, data2Pointer)
	fmt.Println("size:", unsafe.Sizeof(data1), unsafe.Sizeof(data2))

	distance := uintptr(data2Pointer) - uintptr(data1Pointer)
	fmt.Println("distance:", distance)
	fmt.Println("waste:", distance-unsafe.Sizeof(data1))
}
