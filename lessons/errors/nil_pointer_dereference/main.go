package main

import "fmt"

func dererence(pointer *int) {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	_ = *pointer
}

func main() {
	dererence(nil)
}
