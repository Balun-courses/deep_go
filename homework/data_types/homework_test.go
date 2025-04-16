package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"unsafe"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	var result uint32
	pointer := unsafe.Pointer(&number)

	for i := 0; i < 4; i++ {
		b := *(*uint8)(unsafe.Add(pointer, i))
		result = result<<8 | uint32(b)
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
