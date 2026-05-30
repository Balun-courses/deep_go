package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

const ByteSize = 8
const ByteMask = 0xFF

type Number interface {
	uint16 | uint32 | uint64
}

// GetByte получает определенный байт по индексу
func GetByte[T Number](number T, index int) uint8 {
	return uint8(number >> (index * ByteSize) & ByteMask) // сдвигаем на index*8 и отрезаем остальное
}

// ToLittleEndian смена порядка байт Big -> Little Endian.
func ToLittleEndian[T Number](number T) T {
	var result T
	bytesAmount := int(unsafe.Sizeof(number))

	for index := range bytesAmount {
		// Берем index байт из числа
		current := T(GetByte(number, index))

		/*
			Вычисляем смещение для результата в битах - инверсная логика, первый бит окажется последним.
			Используем OR - в нужной нам позиции стоит 0 (null value), OR подставит туда значение.
			Например, для первой итерации для number = 0x12345678:
				- shift = (4 - 1 - 0) * 8 = 24 (самая старшая позиция)
				- current байт = 0x78
				- result = 0x00000000 | 0x78000000 = 0x78000000
		*/

		shift := (bytesAmount - 1 - index) * ByteSize
		result |= current << shift
	}

	return result
}

func TestUint16(t *testing.T) {
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
		})
	}
}

func TestUint32(t *testing.T) {
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

func TestUint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
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
			number: 0x00000000FFFFFFFF,
			result: 0xFFFFFFFF00000000,
		},
		"test case #5": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
		"test case #6": {
			number: 0x1234567890ABCDEF,
			result: 0xEFCDAB9078563412,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
