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
		// can increase capacity
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

	defer allocator.Free()

	pointer1, _ := allocator.Allocate(2)
	pointer2, _ := allocator.Allocate(4)

	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	fmt.Println("address1:", pointer1)
	fmt.Println("address2:", pointer2)
}
