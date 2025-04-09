package main

import "unsafe"

func ToLittleEndian(number uint32) uint32 {
	return ((number & 0xFF000000) >> 24) |
		((number & 0x00FF0000) >> 8) |
		((number & 0x0000FF00) << 8) |
		((number & 0x000000FF) << 24)
}

//задание со звездочкой

type Byte interface {
	~uint32 | ~uint64 | ~uint16
}

func ToLittleEndianGeneric[T Byte](val T) T {
	size := int(unsafe.Sizeof(val))
	var result T
	valPtr := (*[8]byte)(unsafe.Pointer(&val))
	resPtr := (*[8]byte)(unsafe.Pointer(&result))

	for i := 0; i < size; i++ {
		resPtr[i] = valPtr[size-1-i]
	}

	return result
}
