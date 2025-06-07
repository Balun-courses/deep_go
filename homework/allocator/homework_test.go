package main

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	// need to implement
	p := make(map[unsafe.Pointer]*unsafe.Pointer, len(pointers))
	for i := range len(pointers) {
		p[pointers[i]] = &pointers[i]
	}

	var end int
	for i := 0; i < len(memory); i++ {
		if v, ok := p[unsafe.Pointer(&memory[i])]; ok {
			if i == end {
				end++
				for _, ok := p[unsafe.Pointer(&memory[end])]; !ok && memory[end] != 0; _, ok = p[unsafe.Pointer(&memory[end])] {
					end++
				}
				continue
			}
			*v = unsafe.Pointer(&memory[end])
			for t := i; t < len(memory); t++ {
				memory[end] = memory[t]
				memory[t] = 0x00
				end++
				if t+1 == len(memory) {
					i = t + 1
					break
				}
				if _, ok := p[unsafe.Pointer(&memory[t+1])]; ok || memory[t+1] == 0 {
					i = t + 1
					break
				}
			}
		}
	}
}

func TestDefragmentation(t *testing.T) {
	var fragmentedMemory = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0xFF, 0xFF, 0x00,
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
		unsafe.Pointer(&fragmentedMemory[4]),
		unsafe.Pointer(&fragmentedMemory[6]),
		unsafe.Pointer(&fragmentedMemory[7]),
	}

	var defragmentedMemory = []byte{
		0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	Defragment(fragmentedMemory, fragmentedPointers)
	assert.True(t, reflect.DeepEqual(defragmentedMemory, fragmentedMemory))
	assert.True(t, reflect.DeepEqual(defragmentedPointers, fragmentedPointers))
}
