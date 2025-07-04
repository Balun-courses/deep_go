package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// go test -bench=. homework_test.go
// goos: linux
// goarch: amd64
// cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
// BenchmarkSimple-8   	1000000000	         0.8791 ns/op
// BenchmarkUnsafe-8   	749191753	         1.606 ns/op
//
// т.е. простой вариант ToLittleEndian в 2 раза быстрее

// ToLittleEndian - реализация без unsafe
func ToLittleEndian(number uint32) uint32 {
	var result uint32
	for i := 0; i < 4; i++ {
		// отщипываем один байт из number, идя справа налево
		// переносим этот байт в result, заполняя его уже слева направо
		// таким образом, порядок байт меняется на обратный
		curByte := number % 256
		number >>= 8
		result = result<<8 + curByte
	}

	return result
}

// ToLittleEndianUnsafe - реализация c unsafe.Pointer
func ToLittleEndianUnsafe(number uint32) uint32 {
	var result uint32
	inUnsafePointer := unsafe.Add(unsafe.Pointer(&number), 3) // указывает на последний байт number
	outUnsafePointer := unsafe.Pointer(&result)               // указывает на первый байт result
	for i := 0; i < 4; i++ {
		inBytePointer := (*uint8)(inUnsafePointer)
		outBytePointer := (*uint8)(outUnsafePointer)
		*outBytePointer = *inBytePointer // копируем байтик из in в out

		// смещаем указатели: in на байт влево, out - на байт вправо
		inUnsafePointer = unsafe.Add(inUnsafePointer, -1)
		outUnsafePointer = unsafe.Add(outUnsafePointer, 1)
	}

	return result
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

var bench_num = uint32(0x0000FFFF)

func BenchmarkSimple(b *testing.B) {
	for j := 0; j < b.N; j++ {
		_ = ToLittleEndian(bench_num)
	}
}

func BenchmarkUnsafe(b *testing.B) {
	for j := 0; j < b.N; j++ {
		_ = ToLittleEndianUnsafe(bench_num)
	}
}
