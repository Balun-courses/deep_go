package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value1 reflect.Value
	fmt.Println(value1)

	value2 := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(value2)

	value3 := reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(value3)
}
