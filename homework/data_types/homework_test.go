package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type UnsignedNumber interface {
	uint16 | uint32 | uint64 | uint
}

func ToLittleEndian[T UnsignedNumber](number T) T {
	const (
		byteMask  = 0xFF // Маска для первого байта
		byteShift = 8    // Кол-во битов в одном байте
	)

	var (
		maxValue = ^T(0) // Максимальное значение 0xFF...
		res      T
	)

	for remainingBytes := maxValue; remainingBytes > 0; remainingBytes >>= byteShift {
		res <<= byteShift
		res |= number & byteMask

		number >>= byteShift
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
