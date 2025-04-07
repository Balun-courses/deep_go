package main

import "fmt"

const (
	OpenModeIn     = 1 // 0000 0001
	OpenModeOut    = 2 // 0000 0010
	OpenModeAppend = 4 // 0000 0100
	OpenModeBinary = 8 // 0000 1000

	// sugar for prepared masks
	OpenModeInAndOut = OpenModeIn | OpenModeOut // 0000 0001 + 0000 0010 = 0000 0011
)

func Open(filename string, mask int8) {
	if mask&OpenModeIn != 0 {
		fmt.Println("in mode")
	}
	if mask&OpenModeOut != 0 {
		fmt.Println("out mode")
	}
	if mask&OpenModeAppend != 0 {
		fmt.Println("append mode")
	}
	if mask&OpenModeBinary != 0 {
		fmt.Println("binary mode")
	}

	// implementation...
}

func main() {
	Open("data.bin", OpenModeIn|OpenModeBinary) // 0000 1001
	Open("data.bin", OpenModeIn|OpenModeOut)    // 0000 0011
	Open("data.bin", OpenModeAppend)            // 0000 0100
}
