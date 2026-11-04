package main

import (
	"reflect"
	"sort"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// Defragment перемещает все занятые байты (те, на которые ссылаются указатели)
// в начало memory, сохраняя их относительный порядок, обнуляет освободившуюся
// память и обновляет указатели на новые адреса. Всё происходит in-place.
func Defragment(memory []byte, pointers []unsafe.Pointer) {
	if len(memory) == 0 {
		return
	}

	base := uintptr(unsafe.Pointer(&memory[0]))

	// Уникальные адреса занятых байтов в порядке возрастания.
	seen := make(map[uintptr]struct{}, len(pointers))
	occupied := make([]uintptr, 0, len(pointers))
	for _, p := range pointers {
		a := uintptr(p)
		if _, ok := seen[a]; !ok {
			seen[a] = struct{}{}
			occupied = append(occupied, a)
		}
	}
	sort.Slice(occupied, func(i, j int) bool { return occupied[i] < occupied[j] })

	// Адрес -> новый индекс после дефрагментации.
	newIndex := make(map[uintptr]int, len(occupied))

	write := 0
	for _, a := range occupied {
		off := int(a - base)
		// write <= off всегда (обрабатываем адреса по возрастанию),
		// поэтому перезапись не затирает ещё не прочитанные данные.
		memory[write] = memory[off]
		newIndex[a] = write
		write++
	}

	// Освободившуюся память обнуляем.
	for i := write; i < len(memory); i++ {
		memory[i] = 0
	}

	// Перенаправляем указатели на новые адреса.
	for i, p := range pointers {
		pointers[i] = unsafe.Pointer(base + uintptr(newIndex[uintptr(p)]))
	}
}

func TestDefragmentation(t *testing.T) {
	var fragmentedMemory = []byte{
		0xFF, 0x00, 0x00, 0x00,
		0x00, 0xFF, 0x00, 0x00,
		0x00, 0x00, 0xFF, 0x00,
		0x00, 0x00, 0x00, 0xFF,
	}

	var fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[5]),
		unsafe.Pointer(&fragmentedMemory[10]),
		unsafe.Pointer(&fragmentedMemory[15]),
	}

	var defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[1]),
		unsafe.Pointer(&fragmentedMemory[2]),
		unsafe.Pointer(&fragmentedMemory[3]),
	}

	var defragmentedMemory = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	Defragment(fragmentedMemory, fragmentedPointers)
	assert.True(t, reflect.DeepEqual(defragmentedMemory, fragmentedMemory))
	assert.True(t, reflect.DeepEqual(defragmentedPointers, fragmentedPointers))
}

func TestDefragmentationEmptyMemory(t *testing.T) {
	Defragment(nil, nil)
	Defragment([]byte{}, []unsafe.Pointer{})
}

func TestDefragmentationNoPointers(t *testing.T) {
	memory := []byte{1, 2, 3, 4, 5}
	expected := []byte{0, 0, 0, 0, 0}

	Defragment(memory, nil)

	assert.True(t, reflect.DeepEqual(expected, memory))
}

func TestDefragmentationFullyOccupied(t *testing.T) {
	memory := []byte{0x0A, 0x0B, 0x0C, 0x0D}
	pointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[0]),
		unsafe.Pointer(&memory[1]),
		unsafe.Pointer(&memory[2]),
		unsafe.Pointer(&memory[3]),
	}
	expectedMemory := []byte{0x0A, 0x0B, 0x0C, 0x0D}
	expectedPointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[0]),
		unsafe.Pointer(&memory[1]),
		unsafe.Pointer(&memory[2]),
		unsafe.Pointer(&memory[3]),
	}

	Defragment(memory, pointers)

	assert.True(t, reflect.DeepEqual(expectedMemory, memory))
	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestDefragmentationMultiplePointersToSameByte(t *testing.T) {
	memory := []byte{0x00, 0x00, 0x7E, 0x00, 0x00, 0x11, 0x00}
	pointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[5]), // 0x11
		unsafe.Pointer(&memory[2]), // 0x7E
		unsafe.Pointer(&memory[2]), // дубликат указателя на тот же байт
	}
	expectedMemory := []byte{0x7E, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00}
	expectedPointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[1]), // 0x11 был по индексу 5
		unsafe.Pointer(&memory[0]), // 0x7E был по индексу 2
		unsafe.Pointer(&memory[0]), // дубликат
	}

	Defragment(memory, pointers)

	assert.True(t, reflect.DeepEqual(expectedMemory, memory))
	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func TestDefragmentationSingleByte(t *testing.T) {
	memory := []byte{0x42}
	pointers := []unsafe.Pointer{unsafe.Pointer(&memory[0])}

	Defragment(memory, pointers)

	assert.Equal(t, byte(0x42), memory[0])
	assert.Equal(t, unsafe.Pointer(&memory[0]), pointers[0])
}

func TestDefragmentationPreservesOrderAndValues(t *testing.T) {
	// Заняты байты с разными значениями, в т.ч. нулевыми:
	// свободные и занятые различаются только указателями.
	memory := []byte{
		0x00, 0x01, 0x00,
		0x00, 0x00, 0x02,
		0x00, 0x03, 0x00,
		0x04, 0x00, 0x00,
	}
	pointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[9]),  // 0x04
		unsafe.Pointer(&memory[1]),  // 0x01
		unsafe.Pointer(&memory[5]),  // 0x02
		unsafe.Pointer(&memory[7]),  // 0x03
		unsafe.Pointer(&memory[2]),  // 0x00 (занятый нулевой байт)
	}
	expectedMemory := []byte{
		0x01, 0x00, 0x02, 0x03, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	expectedPointers := []unsafe.Pointer{
		unsafe.Pointer(&memory[4]), // 0x04 был по индексу 9
		unsafe.Pointer(&memory[0]), // 0x01 был по индексу 1
		unsafe.Pointer(&memory[2]), // 0x02 был по индексу 5
		unsafe.Pointer(&memory[3]), // 0x03 был по индексу 7
		unsafe.Pointer(&memory[1]), // 0x00 был по индексу 2
	}

	Defragment(memory, pointers)

	assert.True(t, reflect.DeepEqual(expectedMemory, memory))
	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}
