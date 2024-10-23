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
	printAllocs()
	data := make(map[int][128]byte) // try to use 129
	printAllocs()

	for i := 0; i < 1_000_000; i++ {
		data[i] = [128]byte{}
	}

	printAllocs()

	for i := 0; i < 1_000_000; i++ {
		delete(data, i)
	}

	runtime.GC()
	printAllocs()
	runtime.KeepAlive(data)
}
