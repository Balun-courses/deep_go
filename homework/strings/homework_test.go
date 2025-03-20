package main

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type COWBuffer struct {
	data []byte
	refs *int
	mx   *sync.Mutex
}

func finalizer(b *COWBuffer) {
	b.Close()
}

func NewCOWBuffer(data []byte) COWBuffer {
	refs := 1

	b := COWBuffer{
		data: data,
		refs: &refs,
		mx:   &sync.Mutex{},
	}

	runtime.SetFinalizer(&b, finalizer)

	return b
}

func (b *COWBuffer) Clone() COWBuffer {
	b.mx.Lock()
	*b.refs++
	defer b.mx.Unlock()
	return COWBuffer{
		data: b.data,
		refs: b.refs,
		mx:   b.mx,
	}
}

func (b *COWBuffer) Close() {
	b.mx.Lock()
	*b.refs--
	if *b.refs <= 0 {
		b.data = nil
	}
	defer b.mx.Unlock()
}

func (b *COWBuffer) Update(index int, value byte) bool {
	n := len(b.data)
	if index < 0 || index >= n {
		return false
	}

	b.mx.Lock()
	defer b.mx.Unlock()

	// никто больше не ссылается на буффер
	if *b.refs <= 1 {
		b.data[index] = value
		return true
	}

	data := make([]byte, n)
	copy(data, b.data)
	data[index] = value

	// изменяем общий счетчик в любом случае
	*b.refs--

	// создаем рельную отдельную копию
	b.data = data
	refs := 1
	b.refs = &refs
	b.mx = &sync.Mutex{}

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
