package main

import "fmt"

func process(values ...*int) {
	fmt.Println(len(values))
}

func main() {
	process(nil)
	process(nil...)
}
