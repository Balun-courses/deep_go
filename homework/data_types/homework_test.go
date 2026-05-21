package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	size := int(unsafe.Sizeof(number))
	pointer := unsafe.Pointer(&number)

	slice := unsafe.Slice((*byte)(pointer), size)
	reverseSlice := make([]byte, size)
	for i := 0; i < size; i++ {
		reverseSlice[i] = slice[size-1-i]
	}

	result := *(*T)(unsafe.Pointer(&reverseSlice[0]))
	return result
}

func TestСonversion(t *testing.T) {
	testsInt16 := getTestsInt16()
	testsInt32 := getTestsInt32()
	testsInt64 := getTestsInt64()
	runTests(t, testsInt16)
	runTests(t, testsInt32)
	runTests(t, testsInt64)
}

func TestСonversionInt16(t *testing.T) {
	testsInt16 := getTestsInt16()
	runTests(t, testsInt16)
}

func runTests[T uint16 | uint32 | uint64](
	t *testing.T,
	tests map[string]struct {
		number T
		result T
	},
) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func getTestsInt16() map[string]struct {
	number uint16
	result uint16
} {
	return map[string]struct {
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
			number: 0x0102,
			result: 0x0201,
		},
	}
}

func getTestsInt32() map[string]struct {
	number uint32
	result uint32
} {
	return map[string]struct {
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
}

func getTestsInt64() map[string]struct {
	number uint64
	result uint64
} {
	return map[string]struct {
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
			number: 0x0000FFFF0000FFFF,
			result: 0xFFFF0000FFFF0000,
		},
	}
}
