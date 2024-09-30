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

func main() {
	data := make([]byte, 1<<31)

	printAllocs()
	runtime.GC()
	printAllocs()

	fmt.Println(len(data))
}
