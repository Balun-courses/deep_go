package main

import (
	"fmt"
	"testing"
)

var result string

func Concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = "Hello" + " world"
	}
}

func Conversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = string([]byte("Hello world"))
	}
}

func main() {
	b1 := testing.Benchmark(Concat)
	fmt.Println("b1:", b1.AllocsPerOp())       // 0
	fmt.Println("b1:", b1.AllocedBytesPerOp()) // 0

	b2 := testing.Benchmark(Conversion)
	fmt.Println("b2:", b2.AllocsPerOp())       // 1
	fmt.Println("b2:", b2.AllocedBytesPerOp()) // 16
}
