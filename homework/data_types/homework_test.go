package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// ToLittleEndian converts a uint16, uint32, or uint64 value to little-endian format
// by reversing its byte order using unsafe pointer arithmetic.
func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	start := 0
	end := int(unsafe.Sizeof(*new(T)) - 1)
	pointer := unsafe.Pointer(&number)

	for start < end {
		startByte := unsafe.Add(pointer, start)
		endByte := unsafe.Add(pointer, end)

		*(*byte)(startByte), *(*byte)(endByte) = *(*byte)(endByte), *(*byte)(startByte)

		start++
		end--
	}
	return *(*T)(pointer)
}

func TestConversion(t *testing.T) {
	t.Run("uint16 cases", func(t *testing.T) {
		tests := map[string]struct {
			number uint16
			result uint16
		}{
			"all zero":    {number: 0x0000, result: 0x0000},
			"all ones":    {number: 0xFFFF, result: 0xFFFF},
			"alternating": {number: 0x00FF, result: 0xFF00},
			"incremental": {number: 0x1234, result: 0x3412},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				result := ToLittleEndian(test.number)
				assert.Equal(t, test.result, result)
			})
		}
	})

	t.Run("uint32 cases", func(t *testing.T) {
		tests := map[string]struct {
			number uint32
			result uint32
		}{
			"all zero":     {number: 0x00000000, result: 0x00000000},
			"all ones":     {number: 0xFFFFFFFF, result: 0xFFFFFFFF},
			"alternating":  {number: 0x00FF00FF, result: 0xFF00FF00},
			"half full":    {number: 0x0000FFFF, result: 0xFFFF0000},
			"incremental":  {number: 0x01020304, result: 0x04030201},
			"high bit set": {number: 0x80000001, result: 0x01000080},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				result := ToLittleEndian(test.number)
				assert.Equal(t, test.result, result)
			})
		}
	})

	t.Run("uint64 cases", func(t *testing.T) {
		tests := map[string]struct {
			number uint64
			result uint64
		}{
			"all zero":     {number: 0x0000000000000000, result: 0x0000000000000000},
			"all ones":     {number: 0xFFFFFFFFFFFFFFFF, result: 0xFFFFFFFFFFFFFFFF},
			"alternating":  {number: 0x00FF00FF00FF00FF, result: 0xFF00FF00FF00FF00},
			"incremental":  {number: 0x0102030405060708, result: 0x0807060504030201},
			"high bit set": {number: 0x8000000000000001, result: 0x0100000000000080},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				result := ToLittleEndian(test.number)
				assert.Equal(t, test.result, result)
			})
		}
	})
}
