package main

import (
	"fmt"
	"os"
)

func process() {
	defer func() {
		recover()
	}()

	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	os.Exit(1)
}

func main() {
	process()
}
