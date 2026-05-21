package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// Реверс массива байтов через указатель
func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	var littleEndian [8]int8 // т.к. размер должен быть константой, то приходится выделять всегда по-максимуму

	pointer := unsafe.Pointer(&number)

	length := int(unsafe.Sizeof(number))
	for i := 0; i < length; i++ {
		littleEndian[length-1-i] = *(*int8)(unsafe.Add(pointer, i))
	}

	return *(*T)(unsafe.Pointer(&littleEndian))
}

// Реверс через битовые сдвиги (предпочтительный вариант)
func ToLittleEndianByBits[T uint16 | uint32 | uint64](number T) T {
	var length = int(unsafe.Sizeof(number))
	var result T
	for i := 0; i < length; i++ {
		nextByte := byte(number >> (i * 8))
		result |= T(nextByte) << ((length - 1 - i) * 8)
	}
	return result
}

func TestConversion16(t *testing.T) {
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
			number: 0x0102,
			result: 0x0201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)

			result = ToLittleEndianByBits(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion32(t *testing.T) {
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

			result = ToLittleEndianByBits(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion64(t *testing.T) {
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
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)

			result = ToLittleEndianByBits(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
