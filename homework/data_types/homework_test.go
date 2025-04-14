package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// func ToLittleEndian(number uint32) uint32 {
// 	return ((number & 0xFF000000) >> 24) |
// 		((number & 0x00FF0000) >> 8) |
// 		((number & 0x0000FF00) << 8) |
// 		((number & 0x000000FF) << 24)
// }

func TestСonversion32(t *testing.T) {
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
			result := ToLittleEndianAnyUint[uint32](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

// ============== STAR !!!!! ==================//

// Encode big endian to little endian for one of the following types: uint16 | uint32 | uint 64
func ToLittleEndianAnyUint[T uint16 | uint32 | uint64](number T) T {
	size := unsafe.Sizeof(number) //2, 4, 8
	num := uint64(number)
	result := ((num & 0xFF00000000000000) >> 56) |
		((num & 0x00FF000000000000) >> 40) |
		((num & 0x0000FF0000000000) >> 24) |
		((num & 0x000000FF00000000) >> 8) |
		((num & 0x00000000000000FF) << 56) |
		((num & 0x000000000000FF00) << 40) |
		((num & 0x0000000000FF0000) << 24) |
		((num & 0x00000000FF000000) << 8)

	if size < 8 {
		result = result >> (64 - (size * 8))
	}

	return T(result)
}

func TestСonversion16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0000,
			result: 0x0000,
		},
		"test case #2": {
			number: 0xFFFF,
			result: 0xFFFF,
		},
		"test case #3": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #4": {
			number: 0x0F0F,
			result: 0xF0F0,
		},
		"test case #5": {
			number: 0x0102,
			result: 0x2010,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianAnyUint[uint16](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestСonversion64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
		},
		"test case #2": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF00FF00FF,
			result: 0xFF00FF00FF00FF00,
		},
		"test case #4": {
			number: 0x00000000FFFFFFFF,
			result: 0xFFFFFFFF00000000,
		},
		"test case #5": {
			number: 0x0102030405060708,
			result: 0x8070605040302010,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianAnyUint[uint64](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
