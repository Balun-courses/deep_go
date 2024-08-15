package main

import "unsafe"

// small buffer optimization
type SBO struct {
	size  int64
	union [16]byte // 4B[capacity]4B[pointer]
}

func main() {
	var small SBO
	small.size = 10

	var big SBO
	big.size = 1024

	pointer := (**[1024]byte)(unsafe.Pointer(&big.union))
	*(pointer) = new([1024]byte)

	capacity := (*int64)(unsafe.Add(unsafe.Pointer(&big.union), 8))
	*(capacity) = 2048
}
