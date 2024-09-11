package main

import "fmt"

// go build -gcflags='-m' . | grep escape

func main() {
	var value int = 100
	var i interface{} = value
	fmt.Println(i)

	value = 200
	fmt.Println(i)
}
