package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T1 struct {
		a struct{}
		x int64
	}

	fmt.Println(unsafe.Sizeof(T1{}))

	type T2 struct {
		x int64
		a struct{}
	}

	fmt.Println(unsafe.Sizeof(T2{}))

}
