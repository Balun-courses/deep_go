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

	var t1 T1
	fmt.Println("size:", unsafe.Sizeof(t1))
	/*
		fmt.Println("address a:", unsafe.Pointer(&t1.a))
		fmt.Println("address x:", unsafe.Pointer(&t1.x))
	*/

	type T2 struct {
		x int64
		a struct{}
	}

	var t2 T2
	fmt.Println("size:", unsafe.Sizeof(t2))
	/*
		fmt.Println("address a:", unsafe.Pointer(&t2.a))
		fmt.Println("address x:", unsafe.Pointer(&t2.x))
	*/
}
