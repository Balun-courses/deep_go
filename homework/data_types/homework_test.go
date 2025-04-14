package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type uintType interface {
	~uint16 | ~uint32 | ~uint64
}

func ReverseUint16(n uint16) uint16 {
	return (n >> 8) | ((n & 0xff) << 8)
}

func ReverseUint32(n uint32) uint32 {
	return (n >> 24) |
		(((n >> 16) & 0xff) << 8) |
		(((n >> 8) & 0xff) << 16) |
		((n & 0xff) << 24)
}

func ReverseUint64(n uint64) uint64 {
	return (n >> 56) |
		(((n >> 48) & 0xff) << 8) |
		(((n >> 40) & 0xff) << 16) |
		(((n >> 32) & 0xff) << 24) |
		(((n >> 24) & 0xff) << 32) |
		(((n >> 16) & 0xff) << 40) |
		(((n >> 8) & 0xff) << 48) |
		((n & 0xff) << 56)
}

func ReverseBytes[T uintType](n T) T {
	switch any(n).(type) {
	case uint16:
		return T(ReverseUint16(uint16(n)))
	case uint64:
		return T(ReverseUint64(uint64(n)))
	case uint32:
		return T(ReverseUint32(uint32(n)))
	default:
		panic("unreachable")
	}
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
			result := ReverseBytes(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
