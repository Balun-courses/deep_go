package main

import (
	"reflect"
)

func main() {
	pointer1 := reflect.ValueOf(new(int)).Pointer()       // unsafe
	pointer2 := reflect.ValueOf(new(int)).UnsafePointer() // safe

	_ = pointer1
	_ = pointer2
}
