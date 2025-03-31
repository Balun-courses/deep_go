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
		freeObjects: make(map[unsafe.Pointer]struct{}, capacity/objectSize),
		objectSize:  objectSize,
	}

	allocator.resetMemoryState()
	return allocator, nil
}

func (a *PoolAllocator) Allocate() (unsafe.Pointer, error) {
	if len(a.freeObjects) == 0 {
		// can increase capacity
		return nil, errors.New("not enough memory")
	}

	var pointer unsafe.Pointer
	for freePointer := range a.freeObjects {
		pointer = freePointer
		break
	}

	delete(a.freeObjects, pointer)
	return pointer, nil
}

func (a *PoolAllocator) Deallocate(pointer unsafe.Pointer) error {
	if pointer == nil {
		return errors.New("incorrect pointer")
	}

	// potentionally incorrect pointer
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

	defer allocator.Free()

	pointer1, _ := allocator.Allocate()
	pointer2, _ := allocator.Allocate()

	store[int32](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int32](pointer1)
	value2 := load[int32](pointer2)
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	fmt.Println("address1:", pointer1)
	fmt.Println("address2:", pointer2)

	allocator.Deallocate(pointer1)
	allocator.Deallocate(pointer2)
}
