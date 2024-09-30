package main

import (
	"fmt"
)

func main() {
	var x interface{} = 3
	var y interface{} = 3
	fmt.Println(x == y)
}
