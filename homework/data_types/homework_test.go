package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

const (
	b1 uint32 = 0b00000000_00000000_00000000_11111111
	b2 uint32 = 0b00000000_00000000_11111111_00000000
	b3 uint32 = 0b00000000_11111111_00000000_00000000
	b4 uint32 = 0b11111111_00000000_00000000_00000000
)

var bs = [4]uint32{b4, b3, b2, b1}

func ToLittleEndian(number uint32) uint32 {
	var ans uint32 = 0
	for i := 3; i >= 0; i-- {
		n := bs[i] & number
		switch i {
		case 3:
			ans |= (n << 24)
		case 2:
			ans |= n << 8
		case 1:
			ans |= n >> 8
		case 0:
			ans |= n >> 24
		}
	}
	return ans
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
