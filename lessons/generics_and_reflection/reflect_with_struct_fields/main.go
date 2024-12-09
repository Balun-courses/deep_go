package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	Value1 int
	value2 int
}

func main() {
	data := Data{Value1: 100, value2: 200}
	v := reflect.ValueOf(&data).Elem()
	t := v.Type()

	for idx := 0; idx < v.NumField(); idx++ {
		field := v.Field(idx)
		fmt.Println(t.Field(idx).Name, field.CanSet(), field.CanAddr())
	}
}
