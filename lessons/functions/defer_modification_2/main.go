package main

import "fmt"

func Modify(value int) int {
	var result int
	defer func() {
		result += value
	}()

	return value + value
}

func main() {
	fmt.Println(Modify(5))
}
