package main

import (
	"unsafe"
)

type slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func Transform(data []byte) []int {
	sliceData := unsafe.Pointer(&data[0])
	sizeType := int(unsafe.Sizeof(0))
	length := len(data) / sizeType

	var result []int
	resultPtr := (*slice)(unsafe.Pointer(&result))
	resultPtr.data = sliceData
	resultPtr.len = length
	resultPtr.cap = length

	return result
}

func main() {
	data := make([]byte, 800)
	converted := Transform(data)
	_ = converted
}
