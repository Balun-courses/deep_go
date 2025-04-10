package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type customUint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func ToLittleEndian[T customUint](number T) T {
	typeSize := unsafe.Sizeof(number)
	beginPtr := unsafe.Pointer(&number)
	endPtr := unsafe.Add(beginPtr, typeSize-1)

	for range typeSize / 2 {
		leftOctet := (*uint8)(beginPtr)
		rightOctet := (*uint8)(endPtr)
		*leftOctet, *rightOctet = *rightOctet, *leftOctet

		beginPtr = unsafe.Add(beginPtr, 1)
		endPtr = unsafe.Add(endPtr, -1)
	}

	return number
}

func ToLittleEndianBitShift[T customUint](number T) T {
	/*
		number = 4000002010 (uint32)
		11101110 01101011 00101111 11011010

		00000000 00000000 00000000 11101110 ((n >> 24) & 0xff)
		00000000 00000000 01101011 00000000 ((n>>16)&0xff)<<8
		00000000 00101111 00000000 00000000 ((n>>8)&0xff)<<16
		11011010 00000000 00000000 00000000 (n&0xff)<<24
	*/
	typeSize := unsafe.Sizeof(number)
	switch typeSize {
	case 2:
		return (number >> 8) | (number&0xff)<<8
	case 4:
		return ((number >> 24) & 0xff) |
			((number>>16)&0xff)<<8 |
			((number>>8)&0xff)<<16 |
			(number&0xff)<<24
	case 8:
		return ((number >> 56) & 0xff) |
			((number>>48)&0xff)<<8 |
			((number>>40)&0xff)<<16 |
			((number>>32)&0xff)<<24 |
			((number>>24)&0xff)<<32 |
			((number>>16)&0xff)<<40 |
			((number>>8)&0xff)<<48 |
			(number&0xff)<<56
	}
	return number
}

func TestConversion32(t *testing.T) {
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

			result = ToLittleEndianBitShift(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0000,
			result: 0x0000,
		},
		"test case #2": {
			number: 0xFFFF,
			result: 0xFFFF,
		},
		"test case #3": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #4": {
			number: 0x0102,
			result: 0x0201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)

			result = ToLittleEndianBitShift(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x00000000000000000,
			result: 0x00000000000000000,
		},
		"test case #2": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF00FF00FF,
			result: 0xFF00FF00FF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF0000FFFF,
			result: 0xFFFF0000FFFF0000,
		},
		"test case #5": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)

			result = ToLittleEndianBitShift(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func ToLittleEndianBitShiftUint32(number uint32) uint32 {
	return ((number >> 24) & 0xff) |
		((number>>16)&0xff)<<8 |
		((number>>8)&0xff)<<16 |
		(number&0xff)<<24
}

func ToLittleEndianBitShiftUint64(number uint64) uint64 {
	return ((number >> 56) & 0xff) |
		((number>>48)&0xff)<<8 |
		((number>>40)&0xff)<<16 |
		((number>>32)&0xff)<<24 |
		((number>>24)&0xff)<<32 |
		((number>>16)&0xff)<<40 |
		((number>>8)&0xff)<<48 |
		(number&0xff)<<56
}

// BenchmarkConversion-20    	486409120	         2.601 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConversion(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = ToLittleEndian(uint64(0x0102030405060708))
	}
}

// BenchmarkConversionBitShift-20    	481536486	         2.590 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConversionBitShift(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = ToLittleEndianBitShift(uint64(0x0102030405060708))
	}
}

// BenchmarkConversionBitShiftOnly32-20    	1000000000	         0.1715 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConversionBitShiftOnly32(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = ToLittleEndianBitShiftUint32(4000002010)
	}
}

// BenchmarkConversionBitShiftOnly64-20    	1000000000	         0.1705 ns/op	       0 B/op	       0 allocs/op
func BenchmarkConversionBitShiftOnly64(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = ToLittleEndianBitShiftUint64(18040740070709501011)
	}
}
