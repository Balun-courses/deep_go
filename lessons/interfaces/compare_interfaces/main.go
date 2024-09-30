package main

import "fmt"

func main() {
	var lhs interface{} = []int{1, 2, 3}
	var rhs interface{} = []int{1, 2, 3}
	fmt.Println(lhs == rhs)
}
