package main

import (
	"fmt"
	"slices"
)

func sliceToEnd() {
	data := make([]int, 4, 6)

	slice := data[2:]
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 4
}

func sliceMoreThanSize() {
	data := make([]int, 2, 6)

	slice1 := data[1:6]
	_ = slice1
}

func sliceWithIncorrectIndeces() {
	data := make([]int, 2, 6)

	slice2 := data[1:7] // panic
	_ = slice2

	slice3 := data[2:1] // compilation error
	_ = slice3

	left := 2
	right := 1
	slice4 := data[left:right] // panic
	_ = slice4
}

func sliceWithNilSlice() {
	var data []int

	slice := data[:]
	slice = data[0:0]
	slice = data[0:1] // panic
	_ = slice
}

func increaseCapacity() {
	data := make([]int, 0, 10)
	data = data[:10:100] // panic

	slices.Grow(data, 100)
}
