package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type PoolAllocator struct {
	objectPool  []byte
	freeObjects map[unsafe.Pointer]struct{}
	objectSize  int
}

func NewPoolAllocator(capacity int, objectSize int) (PoolAllocator, error) {
	if capacity <= 0 || objectSize <= 0 || capacity%objectSize != 0 {
		return PoolAllocator{}, errors.New("incorrect argumnets")
	}

	allocator := PoolAllocator{
		objectPool:  make([]byte, capacity),
		freeObjects: make(map[unsafe.Pointer]struct{}),
		objectSize:  objectSize,
	}

	allocator.resetMemoryState()
	return allocator, nil
}

func (a *PoolAllocator) Allocate() (unsafe.Pointer, error) {
	if len(a.freeObjects) == 0 {
		return nil, errors.New("not enough memory")
	}

	var pointer unsafe.Pointer
	for freePointer := range a.freeObjects {
		pointer = freePointer
		break
	}

	return pointer, nil
}

func (a *PoolAllocator) Deallocate(pointer unsafe.Pointer) error {
	if pointer == nil {
		return errors.New("incorrect pointer")
	}

	a.freeObjects[pointer] = struct{}{}
	return nil
}

func (a *PoolAllocator) Free() {
	a.resetMemoryState()
}

func (a *PoolAllocator) resetMemoryState() {
	for offset := 0; offset < len(a.objectPool); offset += a.objectSize {
		pointer := unsafe.Pointer(&a.objectPool[offset])
		a.freeObjects[pointer] = struct{}{}
	}
}

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	const KB = 1 << 10
	allocator, err := NewPoolAllocator(KB, 4)
	if err != nil {
		// handling...
	}

	pointer, err := allocator.Allocate()
	if err != nil {
		// handling...
	}

	*(*int32)(pointer) = 100   // 1 way
	store[int32](pointer, 100) // 2 way

	value := *(*int32)(pointer)  // 1 way
	value = load[int32](pointer) // 2 way

	fmt.Println(value)
}
