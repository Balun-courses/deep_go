package main

import "unsafe"

type Data struct {
	id   int32
	data [60]byte
}

// Describe:
// - Alignof
// - Offsetof

var _ uintptr = unsafe.Sizeof(Data{}) - 64
var _ uintptr = 64 - unsafe.Sizeof(Data{})
