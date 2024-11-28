package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value float64 = 3.4

	// неявное преобразование в пустой интерфейс
	fmt.Println("type:", reflect.TypeOf(value))  // func TypeOf(i any) reflect.Type
	fmt.Println("type:", reflect.ValueOf(value)) // func ValueOf(i any) reflect.Value

	equal := reflect.ValueOf(value).Type() == reflect.TypeOf(value)
	fmt.Println("equal:", equal)

	kind := reflect.ValueOf(value).Kind()
	fmt.Println("equal reflect.Float64:", reflect.Float64 == kind)

	value = reflect.ValueOf(value).Float()
	_ = value
}
