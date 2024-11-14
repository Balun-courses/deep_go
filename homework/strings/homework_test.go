package main

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type COWBuffer struct {
	data   []byte
	refs   *int32
	closed bool
	mutex  *sync.RWMutex
}

func NewCOWBuffer(data []byte) COWBuffer {
	return COWBuffer{data, new(int32), false, &sync.RWMutex{}}
}

func (b *COWBuffer) Clone() COWBuffer {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	if b.closed {
		panic("Cannot clone COW string in closed state")
	}

	atomic.AddInt32(b.refs, 1)
	return COWBuffer{b.data, b.refs, false, &sync.RWMutex{}}
}

func (b *COWBuffer) Close() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.closed = true
	*(b.refs)--
}

func (b *COWBuffer) Update(index int, value byte) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if index < 0 || index >= len(b.data) {
		return false
	}

	if b.closed {
		panic("Cannot update COW string in closed state")
	}

	if *b.refs > 1 {
		b.data = append([]byte(nil), b.data...)
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	return *(*string)(unsafe.Pointer(&b.data))
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
