package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
func ToLittleEndianGen[T uint | uint16 | uint32 | uint64](number T) T {
	var res T = 0
	// не придумал как сделать проверку на максимальное значение типа в дженерике
	if number == 0 {
		return number
	}
	ptr := unsafe.Pointer(&number)
	for i := range (int)(unsafe.Sizeof(number)) {
		res = (res << 8) | (T)(*(*uint8)(unsafe.Add(ptr, i)))
	}
	return res
}

func ToLittleEndian(number uint32) uint32 {
	if number == 0 || number == math.MaxUint32 {
		return number
	}
	var res uint32 = 0
	ptr := unsafe.Pointer(&number)
	for i := range (int)(unsafe.Sizeof(number)) {
		res = (res << 8) | (uint32)(*(*uint8)(unsafe.Add(ptr, i)))
	}
	return res // need to implement
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

func TestСonversionGen(t *testing.T) {
	// не придумал как нормально протестировать дженерик
	assert.Equal(t, (uint16)(0x0102), ToLittleEndianGen((uint16)(0x0201)))
	assert.Equal(t, (uint32)(0x01020304), ToLittleEndianGen((uint32)(0x04030201)))
}
