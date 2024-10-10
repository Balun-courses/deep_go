package main

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

const headerSize = 2

type StackAllocator struct {
	data []byte
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
		// can increase header size
		return nil, errors.New("incorrect size")
	}

	previousLength := len(a.data)
	newLength := previousLength + headerSize + size

	if newLength > cap(a.data) {
		// can increase capacity
		return nil, errors.New("not enough memory")
	}

	a.data = a.data[:newLength]
	header := unsafe.Pointer(&a.data[previousLength])
	pointer := unsafe.Pointer(&a.data[previousLength+headerSize])

	*(*int16)(header) = int16(size)
	return pointer, nil
}

func (a *StackAllocator) Deallocate(pointer unsafe.Pointer) error {
	// can deallocate without pointer
	if pointer == nil {
		return errors.New("incorrect pointer")
	}

	header := unsafe.Add(pointer, -headerSize)
	size := *(*int16)(header)

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

	defer allocator.Free()

	pointer1, _ := allocator.Allocate(2)
	defer allocator.Deallocate(pointer1)
	pointer2, _ := allocator.Allocate(4)
	defer allocator.Deallocate(pointer2)

	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	fmt.Println("address1:", pointer1)
	fmt.Println("address2:", pointer2)
}
