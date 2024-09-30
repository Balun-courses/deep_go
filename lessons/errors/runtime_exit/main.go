package main

import (
	"fmt"
	"runtime"
)

func process() {
	defer func() {
		recover()
	}()

	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	runtime.Goexit()
}

func main() {
	process()
}
