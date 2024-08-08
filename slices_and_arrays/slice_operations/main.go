package main

import (
	"fmt"
)

func accessToElement1() {
	data := make([]int, 3)
	fmt.Println(data[4])
}

func accessToElement2() {
	data := make([]int, 3, 6)
	fmt.Println(data[4])
}

func accessToNilSlice() {
	var data []int
	fmt.Println(data[0])
}

func writeToNilSlice() {
	var data []int
	data[0] = 10
}

func appendToNilSlice() {
	var data []int
	data = append(data, 10)
	fmt.Println(data[0])
}

func rangeByNilSlice() {
	var data []int
	for _, value := range data {
		fmt.Println(value)
	}
}

func makeSlice() {
	// _ = make([]int, -5) -> compilation error
	// _ = make([]int, 10, 5) -> compilation error
}

func main() {
}
