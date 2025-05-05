package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

// go test -v homework_test.go
func Defragment(memory []byte, pointers []unsafe.Pointer, dataSize int) error {
	if dataSize < 1 {
		return errors.New("invalid data size")
	}

	if len(memory)%2 != 0 {
		return errors.New("invalid memory size")
	}

	if len(pointers) == 0 || len(memory) == 0 {
		return nil
	}

	currentPointerIdx := 0

	for i := 0; i < len(memory) && currentPointerIdx < len(pointers); {
		if uintptr(unsafe.Pointer(&memory[i])) == uintptr(pointers[currentPointerIdx]) {
			currentPointerIdx++
			i += dataSize
			continue
		}

		isBLockUsed := false

		for j := i; j < i+dataSize && j < len(memory); j++ {
			if memory[j] != 0 {
				isBLockUsed = true
				break
			}
		}

		if isBLockUsed {
			i += dataSize
			continue
		}

		for k := 0; k < dataSize; k++ {
			*(*byte)(unsafe.Add(pointers[currentPointerIdx], k)), memory[i+k] = memory[i+k], *(*byte)(unsafe.Add(pointers[currentPointerIdx], k))
		}
		pointers[currentPointerIdx] = unsafe.Pointer(&memory[i])
		i += dataSize
		currentPointerIdx++
	}

	return nil
}

func TestDefragmentation(t *testing.T) {
	type Data struct {
		fragmentedMemory     []byte
		defragmentedMemory   []byte
		fragmentedPointers   []unsafe.Pointer
		defragmentedPointers []unsafe.Pointer
	}

	data1 := Data{
		fragmentedMemory: []byte{
			0xFF, 0x00, 0x00, 0x00,
			0x00, 0xFF, 0x00, 0x00,
			0x00, 0x00, 0xFF, 0x00,
			0x00, 0x00, 0x00, 0xFF,
		},
		defragmentedMemory: []byte{
			0xFF, 0xFF, 0xFF, 0xFF,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		},
	}

	data1.fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data1.fragmentedMemory[0]),
		unsafe.Pointer(&data1.fragmentedMemory[5]),
		unsafe.Pointer(&data1.fragmentedMemory[10]),
		unsafe.Pointer(&data1.fragmentedMemory[15]),
	}

	data1.defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data1.fragmentedMemory[0]),
		unsafe.Pointer(&data1.fragmentedMemory[1]),
		unsafe.Pointer(&data1.fragmentedMemory[2]),
		unsafe.Pointer(&data1.fragmentedMemory[3]),
	}

	data2 := Data{
		fragmentedMemory: []byte{
			0xFF, 0x00,
			0x00, 0x00,
			0x00, 0xFF,
			0x00, 0x00,
			0x00, 0x00,
			0xFF, 0x00,
			0x00, 0x00,
			0x00, 0xFF,
		},
		defragmentedMemory: []byte{
			0xFF, 0x00,
			0x00, 0xFF,
			0xFF, 0x00,
			0x00, 0xFF,
			0x00, 0x00,
			0x00, 0x00,
			0x00, 0x00,
			0x00, 0x00,
		},
	}

	data2.fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data2.fragmentedMemory[0]),
		unsafe.Pointer(&data2.fragmentedMemory[4]),
		unsafe.Pointer(&data2.fragmentedMemory[10]),
		unsafe.Pointer(&data2.fragmentedMemory[14]),
	}

	data2.defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data2.fragmentedMemory[0]),
		unsafe.Pointer(&data2.fragmentedMemory[2]),
		unsafe.Pointer(&data2.fragmentedMemory[4]),
		unsafe.Pointer(&data2.fragmentedMemory[6]),
	}

	data4 := Data{
		fragmentedMemory: []byte{
			0xFF, 0x00, 0x00, 0x00,
			0x00, 0xFF, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0xFF,
		},
		defragmentedMemory: []byte{
			0xFF, 0x00, 0x00, 0x00,
			0x00, 0xFF, 0x00, 0x00,
			0x00, 0x00, 0x00, 0xFF,
			0x00, 0x00, 0x00, 0x00,
		},
	}

	data4.fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data4.fragmentedMemory[0]),
		unsafe.Pointer(&data4.fragmentedMemory[4]),
		unsafe.Pointer(&data4.fragmentedMemory[12]),
	}

	data4.defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data4.fragmentedMemory[0]),
		unsafe.Pointer(&data4.fragmentedMemory[4]),
		unsafe.Pointer(&data4.fragmentedMemory[8]),
	}

	data8 := Data{
		fragmentedMemory: []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00,
		},
		defragmentedMemory: []byte{
			0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}

	data8.fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data8.fragmentedMemory[8]),
	}

	data8.defragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&data8.fragmentedMemory[0]),
	}

	tests := []struct {
		name    string
		data    Data
		memSize int
	}{
		{
			name:    "1 byte",
			data:    data1,
			memSize: 1,
		},
		{
			name:    "2 bytes",
			data:    data2,
			memSize: 2,
		},
		{
			name:    "4 bytes",
			data:    data4,
			memSize: 4,
		},
		{
			name:    "8 bytes",
			data:    data8,
			memSize: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			_ = Defragment(tt.data.fragmentedMemory, tt.data.fragmentedPointers, tt.memSize)
			assert.True(t, reflect.DeepEqual(tt.data.defragmentedMemory, tt.data.fragmentedMemory))
			assert.True(t, reflect.DeepEqual(tt.data.defragmentedPointers, tt.data.fragmentedPointers))
		})
	}

}
