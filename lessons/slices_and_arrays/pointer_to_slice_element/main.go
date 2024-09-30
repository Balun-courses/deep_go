package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
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
	var numbers = make([]int, 1<<30)
	pointer := FindElement(numbers, 0)
	_ = pointer

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(pointer)
}
