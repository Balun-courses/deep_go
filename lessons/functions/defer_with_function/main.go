package main

import "fmt"

func get() string {
	fmt.Println("1")
	return ""
}

func handle(string) {
	fmt.Println("3")
}

func process() {
	defer handle(get())
	fmt.Println("2")
}

func main() {
	process()
}
