package main

import "unsafe"

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	var pointer1 unsafe.Pointer // not initialized
	var pointer2 unsafe.Pointer // not initialized

	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)

	_ = value1
	_ = value2
}
