package main

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	countRef := 1
	cb := &COWBuffer{
		data: data,
		refs: &countRef,
	}
	runtime.SetFinalizer(cb, func(cb *COWBuffer) {
		cb.Close()
	})
	return *cb
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return *b
}

func (b *COWBuffer) Close() {
	if b.data == nil {
		return
	}
	if b.refs != nil && *b.refs != 0 {
		*b.refs--
	} else {
		b.data = nil
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}
	if *b.refs > 1 {
		*b.refs--
		dataCopy := make([]byte, len(b.data))
		copy(dataCopy, b.data)
		b.data = dataCopy
	}
	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}

func TestCOWBuffer(t *testing.T) {
	data := []byte{'a', 'b', 'c', 'd'}
	buffer := NewCOWBuffer(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	assert.True(t, unsafe.SliceData(data) == unsafe.StringData(buffer.String()))
	assert.True(t, unsafe.StringData(buffer.String()) == unsafe.StringData(copy1.String()))
	assert.True(t, unsafe.StringData(copy1.String()) == unsafe.StringData(copy2.String()))

	assert.True(t, buffer.Update(0, 'g'))
	assert.False(t, buffer.Update(-1, 'g'))
	assert.False(t, buffer.Update(4, 'g'))

	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))

	assert.NotEqual(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))

	copy1.Close()

	previous := copy2.data
	copy2.Update(0, 'f')
	current := copy2.data

	// 1 reference - don't need to copy buffer during update
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))

	copy2.Close()
}
