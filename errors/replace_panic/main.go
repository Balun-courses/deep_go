package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	defer panic(3) // replaced panic(3)
	defer panic(2) // replaced panic(3)
	defer panic(1) // replaced panic(3)
	panic(0)
}
