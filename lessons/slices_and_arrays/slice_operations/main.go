package main

import (
	"fmt"
)

func accessToElement1() {
	data := make([]int, 3)
	fmt.Println(data[4]) // panic
}

func accessToElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4]) // panic
}

func accessToElement3() {
	data := make([]int, 3, 6)
	_ = data[-1] // compilation error
}

func accessToNilSlice1() {
	var data []int
	_ = data[0] // panic
}

func accessToNilSlice2() {
	var data []int
	data[0] = 10 // panic
}

func appendToNilSlice() {
	var data []int
	data = append(data, 10)
}

func rangeByNilSlice() {
	var data []int
	for range data {
	}
}

func makeZeroSlice() {
	data := make([]int, 0)
	fmt.Println(len(data)) // 0
	fmt.Println(cap(data)) // 0
}

func makeSlice() {
	_ = make([]int, -5)    // compilation error
	_ = make([]int, 10, 5) // compilation error

	size := -5
	_ = make([]int, size) // panic

	size = 5
	_ = make([]int, size*2, size) // panic
}
