package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	pointer := unsafe.Pointer(&number)
	pointer2 := unsafe.Add(pointer, 1)
	pointer3 := unsafe.Add(pointer, 2)
	pointer4 := unsafe.Add(pointer, 3)
	u8 := []uint8{
		*(*uint8)(pointer),
		*(*uint8)(pointer2),
		*(*uint8)(pointer3),
		*(*uint8)(pointer4),
	}

	return uint32(u8[3]) | uint32(u8[2])<<8 | uint32(u8[1])<<16 | uint32(u8[0])<<24
}

func TestSerializationProperties(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
