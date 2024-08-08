package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type LinearAllocator struct {
	data []byte
}

func NewLinearAllocator(capacity int) (LinearAllocator, error) {
	if capacity <= 0 {
		return LinearAllocator{}, errors.New("incorrect capacity")
	}

	return LinearAllocator{
		data: make([]byte, 0, capacity),
	}, nil
}

func (a *LinearAllocator) Allocate(size int) (unsafe.Pointer, error) {
	previousLength := len(a.data)
	newLength := previousLength + size

	if newLength > cap(a.data) {
		return nil, errors.New("not enough memory")
	}

	a.data = a.data[:newLength]
	pointer := unsafe.Pointer(&a.data[previousLength])
	return pointer, nil
}

// not supported by this kind of allocator
// func (a *LinearAllocator) Deallocate(pointer unsafe.Pointer) error {}

func (a *LinearAllocator) Free() {
	a.data = a.data[:0]
}

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	const MB = 1 << 20
	allocator, err := NewLinearAllocator(MB)
	if err != nil {
		// handling...
	}

	pointer, err := allocator.Allocate(4)
	if err != nil {
		// handling...
	}

	*(*int32)(pointer) = 100   // 1 way
	store[int32](pointer, 100) // 2 way

	value := *(*int32)(pointer)  // 1 way
	value = load[int32](pointer) // 2 way

	fmt.Println(value)
}
