package main

import (
	"testing"
	"unsafe"
	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	bt0 := byte(number & 0xFF)
	bt1 := byte((number >> 8) & 0xFF)
	bt2 := byte((number >> 16) & 0xFF)
	bt3 := byte((number >> 24) & 0xFF)

	rslt :=  uint32(bt0) | (uint32(bt1) << 8) | (uint32(bt2) << 16) | (uint32(bt3) << 24)
	if !isLittleEndian() {
		rslt =  uint32(bt3) | (uint32(bt2) << 8) | (uint32(bt1) << 16) | (uint32(bt0) << 24)
	}

	return rslt // need to implement
}

func isLittleEndian() bool {
	var checkNum uint16 = 0x0001
	ptr := (*int)(unsafe.Pointer(&checkNum))
	return *ptr == 1
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