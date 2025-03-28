package main

import "unsafe"

// SBO (Small Buffer Optimization)
type SBO struct {
	size  int64
	union [16]byte // 8B[pointer]8B[capacity]
}

func main() {
	var small SBO
	small.size = 10
	small.union = [16]byte{}

	var big SBO
	big.size = 1024
	pointer := (*[2048]byte)(unsafe.Pointer(&big.union))
	*pointer = [2048]byte{}
	capacity := (*int64)(unsafe.Add(unsafe.Pointer(&big.union), 8))
	*capacity = 2048
}
