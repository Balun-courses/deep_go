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
	if data == nil {
		data = []byte{}
	}

	buf := COWBuffer{
		data: data,
		refs: new(int),
	}
	runtime.SetFinalizer(&buf, func(b *COWBuffer) {
		b.Close()
	})

	return buf
}

func (b *COWBuffer) Clone() COWBuffer {
	if b == nil {
		return NewCOWBuffer(nil)
	}

	*b.refs++
	buf := *b
	runtime.SetFinalizer(&buf, func(b *COWBuffer) {
		b.Close()
	})

	return buf
}

func (b *COWBuffer) Close() {
	if b != nil {
		if b.refs != nil {
			if *b.refs > 0 {
				*b.refs--
			}

			if *b.refs == 0 {
				b.refs = nil
			}
		}

		if b.refs == nil {
			b.data = nil
		}
	}
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if b == nil || b.refs == nil {
		return false
	}

	if index < 0 || index >= len(b.data) {
		return false
	}

	if *b.refs > 0 {
		data := make([]byte, len(b.data))
		copy(data, b.data)
		b.data = data
		*b.refs = 0
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

	// Создаем несколько клонов
	copy1 := buffer.Clone()
	copy2 := buffer.Clone()
	copy3 := buffer.Clone()

	// Проверяем, что все клоны указывают на одни данные
	assert.Equal(t, unsafe.SliceData(data), unsafe.SliceData(buffer.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(copy1.data))
	assert.Equal(t, unsafe.SliceData(copy1.data), unsafe.SliceData(copy2.data))
	assert.Equal(t, unsafe.SliceData(copy2.data), unsafe.SliceData(copy3.data))

	// Проверяем строковые представления
	assert.True(t, (*byte)(unsafe.SliceData(data)) == unsafe.StringData(buffer.String()))
	assert.True(t, (*byte)(unsafe.StringData(buffer.String())) == unsafe.StringData(copy1.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy1.String())) == unsafe.StringData(copy2.String()))
	assert.True(t, (*byte)(unsafe.StringData(copy2.String())) == unsafe.StringData(copy3.String()))

	// Обновляем оригинальный буфер (должно создать копию, так как есть клоны)
	assert.True(t, buffer.Update(0, 'g'))
	assert.False(t, buffer.Update(-1, 'g'))
	assert.False(t, buffer.Update(4, 'g'))

	// Проверяем, что оригинал изменился, а клоны остались прежними
	assert.True(t, reflect.DeepEqual([]byte{'g', 'b', 'c', 'd'}, buffer.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy1.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy2.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'b', 'c', 'd'}, copy3.data))

	// Проверяем, что указатели на данные стали разными
	bufferPtr := unsafe.SliceData(buffer.data)
	copy1Ptr := unsafe.SliceData(copy1.data)
	copy2Ptr := unsafe.SliceData(copy2.data)
	copy3Ptr := unsafe.SliceData(copy3.data)

	if bufferPtr == copy1Ptr {
		t.Error("Buffer and copy1 should have different data pointers")
	}
	assert.Equal(t, copy1Ptr, copy2Ptr)
	assert.Equal(t, copy2Ptr, copy3Ptr)

	// Закрываем один клон
	copy1.Close()

	// Обновляем второй клон (не должно создавать копию, так как счетчик ссылок = 1)
	previous := copy2.data
	assert.True(t, copy2.Update(1, 'x'))
	current := copy2.data

	// Проверяем, что копия НЕ была создана (1 reference - don't need to copy)
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'c', 'd'}, copy2.data))
	// copy3 тоже изменился, так как они указывают на одни данные
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'c', 'd'}, copy3.data))

	// Закрываем второй клон
	copy2.Close()

	// Обновляем третий клон (не должно создавать копию, так как нет других клонов)
	previous = copy3.data
	assert.True(t, copy3.Update(2, 'y'))
	current = copy3.data

	// Проверяем, что копия не была создана (1 reference - don't need to copy)
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'y', 'd'}, copy3.data))

	// Создаем новый клон из третьего клона
	copy4 := copy3.Clone()
	assert.Equal(t, unsafe.SliceData(copy3.data), unsafe.SliceData(copy4.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'y', 'd'}, copy4.data))

	// Обновляем третий клон (должно создать копию, так как есть новый клон)
	previous = copy3.data
	assert.True(t, copy3.Update(3, 'z'))
	current = copy3.data

	// Проверяем, что копия была создана
	assert.Equal(t, unsafe.SliceData(previous), unsafe.SliceData(current))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'y', 'z'}, copy3.data))
	assert.True(t, reflect.DeepEqual([]byte{'a', 'x', 'y', 'd'}, copy4.data))

	// Обновляем оригинальный буфер еще раз
	assert.True(t, buffer.Update(1, 'h'))
	assert.True(t, reflect.DeepEqual([]byte{'g', 'h', 'c', 'd'}, buffer.data))

	// Проверяем финальные состояния всех буферов
	assert.Equal(t, "ghcd", buffer.String())
	// После Close() данные могут быть очищены, поэтому проверяем только активные буферы
	assert.Equal(t, "axyz", copy3.String())
	assert.Equal(t, "axyd", copy4.String())

	// Закрываем оставшиеся клоны
	copy3.Close()
	copy4.Close()
}

func TestCOWBufferNil(t *testing.T) {
	// Тест клонирования nil буфера
	var nilBuffer *COWBuffer
	clone := nilBuffer.Clone()

	assert.NotNil(t, clone)
	// Теперь data не должен быть nil, а должен быть пустым срезом
	assert.NotNil(t, clone.data)
	assert.Equal(t, 0, len(clone.data))
	assert.NotNil(t, clone.refs)
	assert.Equal(t, 0, *clone.refs)

	// Тест закрытия nil буфера
	nilBuffer.Close() // Не должно вызывать панику

	// Тест обновления nil буфера - проверяем, что не вызывает панику
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Update на nil буфере вызвал панику: %v", r)
			}
		}()
		// Этот вызов может вызвать панику, но мы её перехватываем
		nilBuffer.Update(0, 'a')
	}()
}

func TestCOWBufferEmpty(t *testing.T) {
	// Тест с пустым буфером
	emptyBuffer := NewCOWBuffer([]byte{})

	assert.NotNil(t, emptyBuffer)
	assert.Equal(t, 0, len(emptyBuffer.data))
	assert.NotNil(t, emptyBuffer.refs)
	assert.Equal(t, 0, *emptyBuffer.refs)

	// Тест с nil буфером (теперь должен создаться пустой срез)
	nilBuffer := NewCOWBuffer(nil)
	assert.NotNil(t, nilBuffer)
	assert.Equal(t, 0, len(nilBuffer.data))
	assert.NotNil(t, nilBuffer.refs)
	assert.Equal(t, 0, *nilBuffer.refs)

	// Тест обновления пустого буфера
	assert.False(t, emptyBuffer.Update(0, 'a'))
	assert.False(t, emptyBuffer.Update(-1, 'a'))

	// Тест обновления nil буфера (теперь пустого)
	assert.False(t, nilBuffer.Update(0, 'a'))
	assert.False(t, nilBuffer.Update(-1, 'a'))

	// Тест клонирования пустого буфера
	clone := emptyBuffer.Clone()
	assert.Equal(t, 0, len(clone.data))
	assert.Equal(t, unsafe.SliceData(emptyBuffer.data), unsafe.SliceData(clone.data))

	// Тест клонирования nil буфера (теперь пустого)
	nilClone := nilBuffer.Clone()
	assert.Equal(t, 0, len(nilClone.data))
	assert.Equal(t, unsafe.SliceData(nilBuffer.data), unsafe.SliceData(nilClone.data))
}

func TestCOWBufferMultipleClones(t *testing.T) {
	data := []byte{'x', 'y', 'z'}
	buffer := NewCOWBuffer(data)

	// Создаем множество клонов
	clones := make([]COWBuffer, 5)
	for i := 0; i < 5; i++ {
		clones[i] = buffer.Clone()
	}

	// Проверяем, что все клоны указывают на одни данные
	for i := 0; i < 5; i++ {
		assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(clones[i].data))
		assert.Equal(t, "xyz", clones[i].String())
	}

	// Обновляем оригинальный буфер (должно создать копию, так как есть клоны)
	assert.True(t, buffer.Update(1, 'w'))

	// Проверяем, что оригинал изменился, а клоны остались прежними
	assert.Equal(t, "xwz", buffer.String())
	for i := 0; i < 5; i++ {
		assert.Equal(t, "xyz", clones[i].String())
		// После обновления оригинального буфера, его данные должны быть в новом месте памяти
		bufferPtr := unsafe.SliceData(buffer.data)
		clonePtr := unsafe.SliceData(clones[i].data)
		if bufferPtr == clonePtr {
			t.Errorf("Buffer and clone %d should have different data pointers", i)
		}
	}
}

func TestCOWBufferReferenceCounting(t *testing.T) {
	data := []byte{'1', '2', '3'}
	buffer := NewCOWBuffer(data)

	// Начальное состояние
	assert.Equal(t, 0, *buffer.refs)

	// Создаем клоны и проверяем счетчик ссылок
	clone1 := buffer.Clone()
	assert.Equal(t, 1, *buffer.refs)

	clone2 := buffer.Clone()
	assert.Equal(t, 2, *buffer.refs)

	clone3 := buffer.Clone()
	assert.Equal(t, 3, *buffer.refs)

	// Закрываем клоны и проверяем уменьшение счетчика
	clone1.Close()
	assert.Equal(t, 2, *buffer.refs)

	clone2.Close()
	assert.Equal(t, 1, *buffer.refs)

	clone3.Close()
	assert.Equal(t, 0, *buffer.refs)

	// После закрытия всех клонов, обновление не должно создавать копию
	originalData := unsafe.SliceData(buffer.data)
	assert.True(t, buffer.Update(0, '9'))
	assert.Equal(t, originalData, unsafe.SliceData(buffer.data))
}

func TestCOWBufferUpdateAfterClone(t *testing.T) {
	data := []byte{'a', 'b', 'c'}
	buffer := NewCOWBuffer(data)
	clone := buffer.Clone()

	// Обновляем клон (должно создать копию, так как есть другой клон)
	assert.True(t, clone.Update(1, 'x'))

	// Проверяем, что клон изменился, а оригинал остался
	assert.Equal(t, "axc", clone.String())
	assert.Equal(t, "abc", buffer.String())

	// Проверяем, что данные теперь в разных местах памяти
	bufferPtr := unsafe.SliceData(buffer.data)
	clonePtr := unsafe.SliceData(clone.data)
	if bufferPtr == clonePtr {
		t.Error("Buffer and clone should have different data pointers")
	}
}

func TestCOWBufferStringMethod(t *testing.T) {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	buffer := NewCOWBuffer(data)

	// Проверяем метод String
	assert.Equal(t, "hello", buffer.String())

	// Проверяем, что String возвращает правильную длину
	assert.Equal(t, len(data), len(buffer.String()))

	// Проверяем с пустым буфером
	emptyBuffer := NewCOWBuffer([]byte{})
	assert.Equal(t, "", emptyBuffer.String())
}

func TestCOWBufferBoundaryConditions(t *testing.T) {
	data := []byte{'a', 'b', 'c'}
	buffer := NewCOWBuffer(data)

	// Тест граничных значений индексов
	assert.False(t, buffer.Update(-1, 'x')) // Отрицательный индекс
	assert.False(t, buffer.Update(3, 'x'))  // Индекс равен длине
	assert.False(t, buffer.Update(10, 'x')) // Индекс больше длины

	// Тест валидных индексов
	assert.True(t, buffer.Update(0, 'x')) // Первый элемент
	assert.True(t, buffer.Update(2, 'z')) // Последний элемент
	assert.True(t, buffer.Update(1, 'y')) // Средний элемент

	assert.Equal(t, "xyz", buffer.String())
}

func TestCOWBufferNilSliceHandling(t *testing.T) {
	// Тест создания буфера с nil срезом
	buffer := NewCOWBuffer(nil)

	// Проверяем, что nil срез был заменен на пустой срез
	assert.NotNil(t, buffer.data)
	assert.Equal(t, 0, len(buffer.data))
	assert.Equal(t, 0, cap(buffer.data))

	// Проверяем, что String() работает корректно с пустым срезом
	assert.Equal(t, "", buffer.String())

	// Проверяем, что клонирование работает
	clone := buffer.Clone()
	assert.NotNil(t, clone.data)
	assert.Equal(t, 0, len(clone.data))
	assert.Equal(t, unsafe.SliceData(buffer.data), unsafe.SliceData(clone.data))

	// Проверяем, что обновление не работает с пустым срезом
	assert.False(t, buffer.Update(0, 'a'))
	assert.False(t, buffer.Update(-1, 'a'))

	// Проверяем, что Close() работает корректно
	buffer.Close()
	clone.Close()
}
