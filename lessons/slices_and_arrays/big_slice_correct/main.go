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

func findSequence(data []byte) []byte {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			partData := make([]byte, 20)
			copy(partData, data[i+2:])
			return partData
		}
	}

	return nil
}

func main() {
	data := make([]byte, 1<<30)
	// let's imagine that data was read from a file

	sequence := findSequence(data)
	_ = sequence // using of sequence later

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(sequence)
}
