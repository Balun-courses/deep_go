package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

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

func TestCOWBufferGenericInt(t *testing.T) {
	data := []int{1, 2, 3, 4}
	buffer := NewCOWBufferGen(data)
	defer buffer.Close()

	copy1 := buffer.Clone()
	copy2 := buffer.Clone()

	assert.True(t, reflect.DeepEqual(buffer.data, copy1.data))
	assert.True(t, reflect.DeepEqual(copy1.data, copy2.data))

	buffer.Update(0, 100)

	assert.True(t, reflect.DeepEqual(buffer.data, []int{100, 2, 3, 4}))
	assert.True(t, reflect.DeepEqual(copy1.data, []int{1, 2, 3, 4}))
	assert.True(t, reflect.DeepEqual(copy2.data, []int{1, 2, 3, 4}))

	copy1.Close()

	old := copy2.data
	copy2.Update(1, 200)

	assert.Equal(t, old, copy2.data)
}
