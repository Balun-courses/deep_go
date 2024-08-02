package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func FindElement(numbers []int, target int) *int {
	for idx := range numbers {
		if numbers[idx] == target {
			return &numbers[idx]
		}
	}

	return nil
}

func main() {
	printAllocs()
	var numbers = make([]int, 10*1024*1024)
	pointer := FindElement(numbers, 0)

	printAllocs()
	// potentially high memory consumption
	_ = pointer

	runtime.GC()
	printAllocs()

	runtime.KeepAlive(pointer)
}
