package main

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Uint interface {
	uint | uint16 | uint32 | uint64
}

func ToLittleEndian[T Uint](number T) T {
	var reversed T

	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, number)

	bufR := bytes.NewReader(buf.Bytes())
	_ = binary.Read(bufR, binary.LittleEndian, &reversed)

	// Забавно, но и в обратном порядке тесты проходят
	//buf := new(bytes.Buffer)
	//_ = binary.Write(buf, binary.LittleEndian, number)
	//
	//bufR := bytes.NewReader(buf.Bytes())
	//_ = binary.Read(bufR, binary.BigEndian, &reversed)

	return reversed
}

func TestConversion(t *testing.T) {
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
