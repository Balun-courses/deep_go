package main

import (
	"fmt"
	"unsafe"
)

func accessToElement1() {
	data := [3]int{1, 2, 3}

	idx := 4
	fmt.Println(data[idx]) // panic

	fmt.Println(data[4]) // compilation error
}

func accessToElement2() {
	data := [3]int{1, 2, 3}

	idx := -1
	fmt.Println(data[idx]) // panic

	fmt.Println(data[-1]) // compilation error
}

func arrayLen() {
	data := [10]int{}
	fmt.Println(len(data)) // 10
}

func capArray() {
	var data [10]int
	fmt.Println(cap(data)) // 10
}

func arraysComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	// except arrays whose element types are incomparable types
	fmt.Println(first == second)
	fmt.Println(first != second)

	//	[<, <=, >, >=]  ->  compilation error
}

func emptyArray() {
	var data [10]byte
	fmt.Println(unsafe.Sizeof(data)) // 10

	//data == nil // compilation error
}

func zeroArray() {
	var data [0]int
	fmt.Println(unsafe.Sizeof(data)) // 0
}

func arrayOfEmptyStructs() {
	var data [10]struct{}
	fmt.Println(unsafe.Sizeof(data)) // 0
}

func negativeArray() {
	var data [-1]int // compilation error
	_ = data
}

func arrayCreation() {
	length1 := 100
	var data1 [length1]int // compilation error
	_ = length1
	_ = data1

	const length2 = 100
	var data2 [length2]int
	_ = data2
}

func makeArray() {
	_ = make([10]int, 10) // compilation error
}

func appendToArray() {
	_ = append([10]int{}, 10) // compilation error
}
