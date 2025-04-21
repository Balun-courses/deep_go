package main

import (
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	ref := 1
	return COWBuffer{
		data: data,
		refs: &ref,
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	if b.refs != nil {
		*b.refs++
	}
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	if b.refs != nil {
		*b.refs--
		if *b.refs < 0 {
			*b.refs = 0
		}
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs > 1 {
		newData := append([]byte(nil), b.data...)
		b.Close()

		ref := 1
		b.data = newData
		b.refs = &ref
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	if len(b.data) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}
