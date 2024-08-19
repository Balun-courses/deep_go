package main

import "fmt"

type Mode struct {
	OpenModeIn     bool
	OpenModeOut    bool
	OpenModeAppend bool
	OpenModeBinary bool
}

func Open(filename string, mode Mode) {
	if mode.OpenModeIn {
		fmt.Println("in mode")
	}
	if mode.OpenModeOut {
		fmt.Println("out mode")
	}
	if mode.OpenModeAppend {
		fmt.Println("append mode")
	}
	if mode.OpenModeBinary {
		fmt.Println("binary mode")
	}

	// implementation...
}

func main() {
	Open("data.bin", Mode{OpenModeIn: true, OpenModeBinary: true})
	Open("data.bin", Mode{OpenModeIn: true, OpenModeOut: true})
	Open("data.bin", Mode{OpenModeAppend: true})
}
