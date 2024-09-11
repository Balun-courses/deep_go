package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int type:", v)
	case string, fmt.Stringer:
		fmt.Println("string type:", v)
	default:
		fmt.Println("unknown type:", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
