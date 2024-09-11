package main

import (
	"fmt"
	"unsafe"
)

type eface struct {
	typ unsafe.Pointer
	val unsafe.Pointer
}

func main() {
	var value int = 100
	var i interface{} = value
	fmt.Println("before:", i)

	obj := (*eface)(unsafe.Pointer(&i))
	*(*int)(obj.val) = 200
	fmt.Println("after:", i)
}
