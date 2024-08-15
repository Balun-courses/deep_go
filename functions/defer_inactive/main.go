package main

import "fmt"

func main() {
	defer fmt.Println("code")
	if false {
		defer fmt.Println("unreacheable code")
	}

	return
	defer fmt.Println("unreacheable code")
}
