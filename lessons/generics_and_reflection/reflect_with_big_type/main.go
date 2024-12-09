package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 10
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())

	x = uint8(v.Uint()) // uint64
	_ = x
}
