package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	return (number&0xff)<<24 |
		((number>>8)&0xff)<<16 |
		((number>>16)&0xff)<<8 |
		(number >> 24)
}

type UnsignedNumbers interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

func SwitchEndianness[T UnsignedNumbers](number T) T {
	var result T
	varByteSize := uint(unsafe.Sizeof(number))

	for pos := range varByteSize {
		result |= ((number >> (pos << 3)) & 0xFF) << ((varByteSize - pos - 1) << 3)
	}
	return result
}

func TestСonversion(t *testing.T) {
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

func TestСonversionGenericUint32(t *testing.T) {
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
			result := SwitchEndianness(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversionGenericUint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1 - zero value": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
		},
		"test case #2 - max value": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3 - alternating bits": {
			number: 0x00FF00FF00FF00FF,
			result: 0xFF00FF00FF00FF00,
		},
		"test case #4 - lower half filled": {
			number: 0x00000000FFFFFFFF,
			result: 0xFFFFFFFF00000000,
		},
		"test case #5 - custom pattern #1": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
		"test case #6 - custom pattern #2": {
			number: 0xAABBCCDDEEFF1122,
			result: 0x2211FFEEDDCCBBAA,
		},
		"test case #7 - single byte shift": {
			number: 0x00000000000000FF,
			result: 0xFF00000000000000,
		},
		"test case #8 - two byte shift": {
			number: 0x000000000000FFFF,
			result: 0xFFFF000000000000,
		},
		"test case #9 - four byte shift": {
			number: 0x00000000FFFFFFFF,
			result: 0xFFFFFFFF00000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := SwitchEndianness(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

