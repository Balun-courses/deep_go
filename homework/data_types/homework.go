package main

import (
	"unsafe"
)

type union interface {
	uint16 | uint32 | uint64
}

const B_SIZE = 8

func ToLittleEndian[T union](number T) T {

	var res T
	for i := 0; i < int(unsafe.Sizeof(number)); i++ {
		res = (res << B_SIZE) | (number & 0xFF)
		number >>= B_SIZE
	}

	return res
}
