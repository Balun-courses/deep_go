package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	shift0 := number
	shift8 := number >> 8
	shift16 := number >> 16
	shift24 := number >> 24

	pointerToFirstByteOfShift0 := (*uint8)(unsafe.Pointer(&shift0))
	pointerToFirstByteOfShift8 := (*uint8)(unsafe.Pointer(&shift8))
	pointerToFirstByteOfShift16 := (*uint8)(unsafe.Pointer(&shift16))
	pointerToFirstByteOfShift24 := (*uint8)(unsafe.Pointer(&shift24))

	var res uint32 = 0

	res += uint32(*pointerToFirstByteOfShift0)
	res = res << 8
	res += uint32(*pointerToFirstByteOfShift8)
	res = res << 8
	res += uint32(*pointerToFirstByteOfShift16)
	res = res << 8
	res += uint32(*pointerToFirstByteOfShift24)

	return res
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
