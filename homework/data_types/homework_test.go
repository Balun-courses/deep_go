package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T uint16 | uint32 | uint64 | uint](number T) T {
	var (
		size = unsafe.Sizeof(number)
		ptr1 = unsafe.Pointer(&number)
		ptr2 = unsafe.Add(unsafe.Pointer(&number), size-1)
	)

	for range int(size) / 2 {
		*(*byte)(ptr1), *(*byte)(ptr2) = *(*byte)(ptr2), *(*byte)(ptr1)

		ptr1 = unsafe.Add(ptr1, 1)
		ptr2 = unsafe.Add(ptr2, -1)
	}

	return number
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
