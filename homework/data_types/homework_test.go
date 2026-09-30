package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./homework/data_types/

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

func testConversion[T endianTypes](t *testing.T, table map[string]struct{ number, result T }) {
	for name, test := range table {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, ToLittleEndian(test.number))
		})
	}
}

func TestConversionUint16(t *testing.T) {
	testConversion(t, map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {number: 0x0000, result: 0x0000},
		"test case #2": {number: 0xFFFF, result: 0xFFFF},
		"test case #3": {number: 0x0102, result: 0x0201},
	})
}

func TestConversionUint64(t *testing.T) {
	testConversion(t, map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {number: 0x0000000000000000, result: 0x0000000000000000},
		"test case #2": {number: 0xFFFFFFFFFFFFFFFF, result: 0xFFFFFFFFFFFFFFFF},
		"test case #3": {number: 0x0102030405060708, result: 0x0807060504030201},
		"test case #4": {number: 0x000000000000FFFF, result: 0xFFFF000000000000},
	})
}
