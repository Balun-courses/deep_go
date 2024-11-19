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
	fmt.Println(i)

	value = 200
	fmt.Println(i)

	obj := (*eface)(unsafe.Pointer(&i))
	println("&value:", &value)
	println("obj.val:", obj.val)

}
