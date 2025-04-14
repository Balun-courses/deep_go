package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func findSequence(data string) string {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\n' && data[i+1] == '\t' {
			return data[i+2 : i+22]
		}
	}

	return ""
}

func main() {
	data := make([]byte, 1<<30)
	// let's imagine that data was read from a file

	str := unsafe.String(unsafe.SliceData(data), len(data))
	sequence := findSequence(str)
	_ = sequence // using of sequence later

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}
