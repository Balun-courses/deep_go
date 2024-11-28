package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value float64 = 3.14
	v := reflect.ValueOf(value) // copy of value
	//v.SetFloat(2.7)
	//v.Addr()

	fmt.Println("settability of v:", v.CanSet())
}
