package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// go test -v homework_test.go

func ToLittleEndianForInteger[T constraints.Integer](number T) T {
	result := number
	pResult := unsafe.Pointer(&result)

	n := int(unsafe.Sizeof(result))
	for offset := 0; offset < n/2; offset++ {
		lp := (*uint8)(unsafe.Add(pResult, offset))
		rp := (*uint8)(unsafe.Add(pResult, n-offset-1))
		tmp := *lp
		*lp = *rp
		*rp = tmp
	}

	return result
}

func ToLittleEndian(number uint32) uint32 {
	return ToLittleEndianForInteger[uint32](number)
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
