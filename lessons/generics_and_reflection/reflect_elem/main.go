package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value float64 = 3.14
	v := reflect.ValueOf(&value) // copy of value

	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	fmt.Println("addresability of v:", v.CanAddr())

	fmt.Println("type of v.Elem():", v.Elem().Type())             // dereference
	fmt.Println("settability of v.Elem():", v.Elem().CanSet())    // dereference
	fmt.Println("addresability of v.Elem():", v.Elem().CanAddr()) // dereference

	v.Elem().SetFloat(2.7)

	fmt.Println(v.Elem().Addr())
	fmt.Println("value:", value)
	fmt.Println("address:", &value)
}
