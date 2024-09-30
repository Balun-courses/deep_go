package main

import "fmt"

func main() {
	var value int = 100
	var i interface{} = value
	fmt.Println(i)

	value = 200
	fmt.Println(i)
}
