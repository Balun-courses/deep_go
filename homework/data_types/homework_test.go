package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	return ((number & 0xFF) << 24) |
		(((number >> 8) & 0xFF) << 16) |
		(((number >> 16) & 0xFF) << 8) |
		((number >> 24) & 0xFF)
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

func ToLittleEndianGeneric[T uint16 | uint32 | uint64](number T) T {
	intLen := unsafe.Sizeof(number)

	if intLen == 2 {
		return ((number & 0xFF) << 8) |
			((number >> 8) & 0xFF)
	}
	if intLen == 4 {
		return ((number & 0xFF) << 24) |
			(((number >> 8) & 0xFF) << 16) |
			(((number >> 16) & 0xFF) << 8) |
			((number >> 24) & 0xFF)
	}
	if intLen == 8 {
		return ((number & 0xFF) << 56) |
			(((number >> 8) & 0xFF) << 48) |
			(((number >> 16) & 0xFF) << 40) |
			(((number >> 24) & 0xFF) << 32) |
			(((number >> 32) & 0xFF) << 24) |
			(((number >> 40) & 0xFF) << 16) |
			(((number >> 48) & 0xFF) << 8) |
			((number >> 56) & 0xFF)
	}

	return 0
}
