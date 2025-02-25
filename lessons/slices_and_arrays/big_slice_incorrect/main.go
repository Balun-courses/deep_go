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
		if data[i] == 0xFF && data[i+1] == 0xEE {
			return data[i : i+20]
		}
	}

	return nil
}

func processBigData() {
	var data []byte
	// let's imagine that data was read from a file

	sequence := findSequence(data)
	_ = sequence // using of sequence later
}
