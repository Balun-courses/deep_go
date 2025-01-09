package main

import "fmt"

func main() {
	var f func()
	defer f()

	f = func() {
		fmt.Println(true)
	}
}
