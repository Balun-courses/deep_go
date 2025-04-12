package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type uints interface {
	~uint16 | ~uint32 | ~uint64
}

func ToLittleEndian[T uints](number T) T {
	// Simple hardcoded solution for uint32
	// return (number>>24)&0xFF | (number>>8)&0xFF00 | (number<<8)&0xFF0000 | (number<<24)&0xFF000000

	// Dynamic for generics solution
	var result T
	var size = int(unsafe.Sizeof(T(number)))
	var bytePos = size - 1

	for i := 0; i < size; i += 1 {
		mask := T(0xFF) << (i * 8)
		result |= ((number & mask) >> (i * 8)) << (bytePos * 8)
		bytePos -= 1
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
