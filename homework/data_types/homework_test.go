package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	// побайтовый сдвиг, старший становится младшим младший - старшим
	// используем маску 0xFF00000 и операцию сложения для выбора байта
	return ((number & 0xFF000000) >> 24) |
		((number & 0x00FF0000) >> 8) |
		((number & 0x0000FF00) << 8) |
		((number & 0x000000FF) << 24)
}

func ToLittleEndianGeneric[T constraints.Unsigned](number T) T {
	// Так как uint8, uint16 не может быть сдвинут на 24, то конвертируем в больший тип,
	// зетем приводим обратно к нужному типу
	// выдаст ошибку на uint64, чтобы сработало нужно увеличить количество сдвигов
	x := uint32(number)
	x = ((x & 0xFF000000) >> 24) |
		((x & 0x00FF0000) >> 8) |
		((x & 0x0000FF00) << 8) |
		((x & 0x000000FF) << 24)

	return T(x)
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
			result := ToLittleEndianGeneric(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
