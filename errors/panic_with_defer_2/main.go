package main

import "fmt"

func process() {
	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	panic("error")
}

func main() {
	process()
}
