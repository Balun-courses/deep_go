package main

import (
	"reflect"
	"runtime"
	"slices"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	refs := 1

	buf := COWBuffer{
		data: data,
		refs: &refs,
	}

	runtime.SetFinalizer(&buf, (*COWBuffer).Close)

	return buf
}

func (b *COWBuffer) Clone() COWBuffer {
	buf := COWBuffer{
		data: b.data,
		refs: b.refs,
	}

	*buf.refs += 1

	return buf
}

func (b *COWBuffer) Close() {
	if b == nil || *b.refs == 0 {
		return
	}

	*b.refs -= 1
	b.data = nil
	runtime.SetFinalizer(b, nil)
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index >= len(b.data) || index < 0 {
		return false
	}

	// На буфер никто не ссылается - не копируем
	if b.refs == nil || *b.refs == 1 {
		b.data[index] = value

		*b.refs = +1
		return true
	}

	// На буффер ссылаются
	// Уменьшаем старый счетчик
	*b.refs -= 1

	refs := 1
	// Делаем новую копию с refs=1
	clone := COWBuffer{
		data: slices.Clone(b.data),
		refs: &refs,
	}
	clone.data[index] = value

	b.data = clone.data
	b.refs = clone.refs

	return true
}

func (b *COWBuffer) String() string {
	data := unsafe.SliceData(b.data)
	return unsafe.String(data, len(b.data))
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

	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(copy1.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy1.String())) == unsafe.StringData(copy2.String()))

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
