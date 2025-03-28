package main

import (
	"fmt"
	"unsafe"
)

type data struct {
	aaa bool
	bbb int32
	ccc bool
}

func main() {
	d := data{}
	fmt.Println(unsafe.Offsetof(d.ccc))

	fmt.Println(unsafe.Alignof(d))
}
