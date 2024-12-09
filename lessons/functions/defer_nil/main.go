package main

import "fmt"

func main() {
	defer fmt.Println("reachable 1")
	var f func()
	defer f()

	fmt.Println("reachable 2")
}
