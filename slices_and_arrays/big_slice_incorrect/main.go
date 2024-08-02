package main

import (
	"os"
)

func FindData(filename string) []byte {
	data, _ := os.ReadFile(filename)

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0xFF && data[i+1] == 0x00 {
			return data[i : i+20] // capacity leak
		}
	}

	return nil
}

func main() {
	_ = FindData("data.bin")
	// potentially high memory consumption
}
