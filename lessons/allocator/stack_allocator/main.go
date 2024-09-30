package main

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

const headerSize = 2

type StackAllocator struct {
	data     []byte
	lastSize int
}

func NewStackAllocator(capacity int) (StackAllocator, error) {
	if capacity <= 0 {
		return StackAllocator{}, errors.New("incorrect capacity")
	}

	return StackAllocator{
		data: make([]byte, 0, capacity),
	}, nil
}

func (a *StackAllocator) Allocate(size int) (unsafe.Pointer, error) {
	if size > math.MaxInt16 {
		return nil, errors.New("incorrect size")
	}

	previousLength := len(a.data)
	newLength := previousLength + headerSize + size

	if newLength > cap(a.data) {
		return nil, errors.New("not enough memory")
	}

	a.data = a.data[:newLength]

	header := unsafe.Pointer(&a.data[previousLength])
	*(*int16)(header) = int16(size)

	pointer := unsafe.Pointer(&a.data[previousLength+headerSize])
	return pointer, nil
}

func (a *StackAllocator) Deallocate(pointer unsafe.Pointer) error {
	if pointer == nil {
		return errors.New("incorrect pointer")
	}

	header := uintptr(pointer) - headerSize
	size := *(*int16)(unsafe.Pointer(header))

	previousLength := len(a.data)
	newLength := previousLength - headerSize - int(size)

	a.data = a.data[:newLength]
	return nil
}

func (a *StackAllocator) Free() {
	a.data = a.data[:0]
}

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	const KB = 1 << 10
	allocator, err := NewStackAllocator(KB)
	if err != nil {
		// handling...
	}

	pointer, err := allocator.Allocate(8)
	if err != nil {
		// handling...
	}

	*(*int64)(pointer) = 100   // 1 way
	store[int64](pointer, 100) // 2 way

	value := *(*int64)(pointer)  // 1 way
	value = load[int64](pointer) // 2 way

	fmt.Println(value)
}
