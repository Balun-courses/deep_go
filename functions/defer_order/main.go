package main

import "fmt"

func process() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func main() {
	process()
}
