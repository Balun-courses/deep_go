package main

import "fmt"

// go build -gcflags='-m' . | grep escape

func main() {
	var intValue = 100
	var boolValue = true
	var strValue = "hello"
	fmt.Println(intValue, boolValue, strValue)
}
