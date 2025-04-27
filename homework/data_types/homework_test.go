package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	var res uint32 = 0
	for i := 3; i >= 0; i-- {
		// take the last byte
		lastByteValue := number % uint32(math.Pow(2, float64(8)))

		// put it to the proper place in the result
		res += lastByteValue * uint32(math.Pow(2, float64(8*i)))

		// remove the last byte
		number -= lastByteValue
		number /= uint32(math.Pow(2, float64(8)))
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
