package main

import (
	"fmt"
	"testing"
)

var result string
var buffer []byte = make([]byte, 33)

func Concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = string(buffer) + string(buffer)
	}
}

func main() {
	b := testing.Benchmark(Concat)
	fmt.Println(b.AllocsPerOp())       // 3
	fmt.Println(b.AllocedBytesPerOp()) // 176

	// alocated 176 bytes = 48 (not 33) + 48 (not 33) + 80 (not 33 + 33 = 66)
	// waste = 15 + 15 + 14 = 44 bytes (25% waste)
}
