package main

import (
	"reflect"
	"unsafe"
)

func main() {
	pointer1 := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer())) // safe
	pointer2 := reflect.ValueOf(new(int)).Pointer()                         // unsafe
	pointer3 := reflect.ValueOf(new(int)).UnsafePointer()                   // safe

	_ = pointer1
	_ = pointer2
	_ = pointer3
}
