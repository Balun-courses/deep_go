package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	writeIndex := 0
	pointerMap := make(map[uintptr]int)

	// создаём карту старых указателей → их индексы в `pointers`
	for i, p := range pointers {
		pointerMap[uintptr(p)] = i
	}

	// перемещаем занятые байты (0xFF) в начало памяти
	for readIndex := 0; readIndex < len(memory); readIndex++ {
		if memory[readIndex] == 0xFF {
			if writeIndex != readIndex {
				memory[writeIndex] = memory[readIndex]
				memory[readIndex] = 0x00
			}
			// если этот байт был одним из указанных в pointers — обновим его
			if i, ok := pointerMap[uintptr(unsafe.Pointer(&memory[readIndex]))]; ok {
				pointers[i] = unsafe.Pointer(&memory[writeIndex])
			}
			writeIndex++
		}
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
