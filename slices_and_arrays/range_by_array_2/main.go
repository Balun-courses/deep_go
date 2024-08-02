package main

import "fmt"

func main() {
	data := [...]int{3, 2, 1, 0}
	pointers := [4]*int{}

	// also actual for slices
	for idx, value := range data {
		pointers[idx] = &value
		//fmt.Println(uintptr(unsafe.Pointer(&value)))
	}

	for _, pointer := range pointers {
		fmt.Println(pointer)
		*pointer = 10
	}

	fmt.Println(data)
}
