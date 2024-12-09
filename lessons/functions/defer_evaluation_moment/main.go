package main

import "fmt"

func main() {
	var f = func() {
		fmt.Println(false)
	}

	defer f()

	f = func() {
		fmt.Println(true)
	}
}
