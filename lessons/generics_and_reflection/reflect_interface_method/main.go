package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value float32 = 3.14
	v := reflect.ValueOf(value)
	iValue := v.Interface()
	fmt.Println(iValue.(float32))
}
