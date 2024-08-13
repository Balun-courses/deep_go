package main

import (
	"fmt"
	"unsafe"
)

type data1 struct {
	aaa bool
	bbb [987]byte // TOD0
	ccc bool
}

func main() {
	d := data1{}
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Alignof(d))
}
