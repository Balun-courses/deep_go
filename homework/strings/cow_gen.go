package main

import (
	"unsafe"
)

type COWBufferGen[T any] struct {
	data []T
	refs *int
}

func NewCOWBufferGen[T any](data []T) COWBufferGen[T] {
	ref := 1
	return COWBufferGen[T]{
		data: data,
		refs: &ref,
	}
}

func (b *COWBufferGen[T]) Clone() COWBufferGen[T] {
	if b.refs != nil {
		*b.refs++
	}
	return COWBufferGen[T]{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBufferGen[T]) Close() {
	if b.refs != nil {
		*b.refs--
		if *b.refs < 0 {
			*b.refs = 0
		}
	}
}

func (b *COWBufferGen[T]) Update(index int, value T) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}
	if *b.refs > 1 {
		newData := append([]T(nil), b.data...)
		b.Close()
		ref := 1
		b.data = newData
		b.refs = &ref
	}
	b.data[index] = value
	return true
}

func (b *COWBufferGen[T]) String() string {
	if _, ok := any(*b).(COWBufferGen[byte]); !ok {
		panic("String() is only valid for COWBuffer[byte]")
	}

	ptr := unsafe.Pointer(unsafe.SliceData(b.data))
	return unsafe.String((*byte)(ptr), len(b.data))
}
