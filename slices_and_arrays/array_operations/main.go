package main

import (
	"fmt"
	"unsafe"
)

func accessToElement() {
	data := [3]int{1, 2, 3}

	idx := 4
	fmt.Println(data[idx])

	// fmt.Println(data[4]) -> compilation error
}

func arrayLength() {
	data := [10]int{}
	fmt.Println(len(data))
}

func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	//	[<, <=, >, >=]  ->  compilation error
}

func nilArray() {
	var data [10]int
	fmt.Println(unsafe.Sizeof(data))

	// data == nil // -> compilation error
}

func emptyArray() {
	var data [0]int
	fmt.Println(unsafe.Sizeof(data))
}

func capArray() {
	/*
		var data [10]int
		_ = cap(data) // -> compilation error
	*/
}

func nonConstantLength() {
	/*
		length := 100
		var data [length]int // -> compilation error
		_ = data
	*/
}

func negativeArray() {
	// var data [-1]int // -> compilation error
}

func makeArray() {
	// _ = make([10]int, 10) // -> compilation error
}

func appendToArray() {
	// _ = append([10]int{}, 10) // -> compilation error
}

func main() {
}
