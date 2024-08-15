package main

import "fmt"

func process1() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
}

func process2() {
	defer func() {
		defer fmt.Println(4)
		defer fmt.Println(3)
	}()

	defer func() {
		defer fmt.Println(2)
		defer fmt.Println(1)
	}()
}

func main() {
	process1()
}
