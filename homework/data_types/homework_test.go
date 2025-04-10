package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	const mask = 0xFF                  // initialize bit mask for one byte
	size := int(unsafe.Sizeof(number)) // number of bytes
	var res T                          // initialize result of type T
	for i := 0; i < size; i++ {        // iterate over bytes from higher order to lower
		res <<= 8            // take all the bytes and make them higher by one order, leaving zero at the lowest byte
		res |= number & mask // take the lowest byte from number and put in into the lowest byte of result
		number = number >> 8 // remove the lowest byte from number, shift right and put the next highest byte to the lowest position
	}
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
