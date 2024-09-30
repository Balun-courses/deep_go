package main

import "fmt"

func main() {
	data := []int{0, 1, 2, 3}
	lookup := make(map[int]*int, 4)

	for idx, value := range data {
		lookup[idx] = &value
		//fmt.Println(uintptr(unsafe.Pointer(&value)))
	}

	for key, value := range lookup {
		fmt.Println(key, *value)
	}
}
