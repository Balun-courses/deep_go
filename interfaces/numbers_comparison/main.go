package main

import (
	"fmt"
	"io"
)

type Interface interface {
	String() string

	interface {
		Clone() Interface
	}

	io.ReadWriter
}

func main() {
	var x interface{} = 3
	var y interface{} = 3
	fmt.Println(x == y)
}
