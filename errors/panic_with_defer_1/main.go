package main

import "fmt"

func processV1() {
	defer func() {
		recover()
	}()

	fmt.Println("V1: open file")
	panic("error")
	fmt.Println("V1: close file")
}

func processV2() {
	defer func() {
		recover()
	}()

	fmt.Println("V2: open file")
	defer fmt.Println("V2: close file")

	panic("error")
}

func main() {
	processV1()
	processV2()
}
