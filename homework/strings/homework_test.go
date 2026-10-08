package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// COWBuffer - copy-on-write буфер.
// data совместно используется всеми копиями буфера,
// refs - разделяемый счетчик ссылок на данные,
// nil в refs означает, что копия буфера закрыта (Close).
type COWBuffer struct {
	data []byte
	refs *int
}

// NewCOWBuffer создает новый буфер с разделяемым счетчиком ссылок.
func NewCOWBuffer(data []byte) COWBuffer {
	refs := 1
	return COWBuffer{data: data, refs: &refs}
}

// Clone создает новую копию буфера: данные остаются общими,
// счетчик ссылок увеличивается.
func (b *COWBuffer) Clone() COWBuffer {
	if b.refs == nil {
		// клонирование закрытого буфера - возвращаем независимый пустой буфер
		return NewCOWBuffer(b.data)
	}
	*b.refs++
	return COWBuffer{data: b.data, refs: b.refs}
}

// Close перестает использовать копию буфера.
// Повторный вызов Close (и любой последующий) отрабатывает корректно -
// это защита от программиста, забывшего, что Close уже был вызван
// (например, явно и затем через defer).
func (b *COWBuffer) Close() {
	if b.refs == nil {
		return
	}
	*b.refs--
	b.data = nil
	b.refs = nil
}

// Update изменяет байт по индексу index.
// Если на данные ссылается кто-то еще (refs > 1), сначала создается
// новая копия данных, и только затем производится изменение.
// Если ссылок больше нет - данные изменяются на месте, без копирования.
func (b *COWBuffer) Update(index int, value byte) bool {
	if b.refs == nil || index < 0 || index >= len(b.data) {
		return false
	}
	if *b.refs > 1 {
		*b.refs-- // открепляемся от общего буфера
		*b = NewCOWBuffer(append([]byte(nil), b.data...))
	}
	b.data[index] = value
	return true
}

// String конвертирует буфер в строку без копирования
// и дополнительного выделения памяти.
func (b *COWBuffer) String() string {
	if len(b.data) == 0 {
		return ""
	}
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

// Множественный вызов Close() должен отрабатывать корректно:
// например, программист вызвал Close() явно и еще раз через defer.
func TestCOWBufferMultipleClose(t *testing.T) {
	data := []byte{'x', 'y', 'z'}
	buffer := NewCOWBuffer(data)
	clone := buffer.Clone()
	defer clone.Close()

	buffer.Close()
	buffer.Close() // повторный вызов Close - безопасен
	buffer.Close()

	// clone по-прежнему владеет данными
	assert.True(t, clone.Update(2, 'w'))
	assert.True(t, reflect.DeepEqual([]byte{'x', 'y', 'w'}, clone.data))

	clone.Close()
	clone.Close() // и еще раз после явного вызова
}

// Использование буфера после Close не должно приводить к панике.
func TestCOWBufferUsageAfterClose(t *testing.T) {
	buffer := NewCOWBuffer([]byte("abc"))
	buffer.Close()

	assert.False(t, buffer.Update(0, 'z'))
	assert.Empty(t, buffer.String())
}

// Проверка границ индекса при Update.
func TestCOWBufferUpdateBounds(t *testing.T) {
	buffer := NewCOWBuffer([]byte("abc"))
	defer buffer.Close()

	assert.True(t, buffer.Update(2, 'C')) // последний валидный индекс
	assert.False(t, buffer.Update(3, 'd')) // ровно за границей
	assert.False(t, buffer.Update(-1, 'd'))

	assert.Equal(t, "abC", buffer.String())
}

// После copy-on-write String() по-прежнему указывает на собственные
// данные буфера (без копирования).
func TestCOWBufferStringZeroCopyAfterUpdate(t *testing.T) {
	buffer := NewCOWBuffer([]byte("hello"))
	defer buffer.Close()

	clone := buffer.Clone()
	defer clone.Close()

	assert.True(t, buffer.Update(0, 'H'))

	assert.True(t, unsafe.StringData(buffer.String()) == unsafe.SliceData(buffer.data))
	assert.Equal(t, "Hello", buffer.String())
	assert.Equal(t, "hello", clone.String())
}

// Пустой буфер.
func TestCOWBufferEmpty(t *testing.T) {
	buffer := NewCOWBuffer(nil)
	defer buffer.Close()

	assert.Empty(t, buffer.String())
	assert.False(t, buffer.Update(0, 'a'))

	clone := buffer.Clone()
	defer clone.Close()
	assert.Empty(t, clone.String())
}
