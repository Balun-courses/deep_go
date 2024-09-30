package main

import "fmt"

func Modify(value int) (result int) {
	defer func() {
		result += value
	}()

	return value + value
}

func main() {
	fmt.Println(Modify(5))
}
