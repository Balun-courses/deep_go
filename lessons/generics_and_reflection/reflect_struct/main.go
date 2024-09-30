package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	Count int
	Title string
}

func (d Data) Do() {}

func main() {
	data := Data{}

	tData := reflect.TypeOf(data)
	fmt.Println("Kind:", tData.Kind())
	fmt.Println("PkgPath:", tData.PkgPath())

	fmt.Println("NumField:", tData.NumField())
	fmt.Println("Field(0):", tData.Field(0).Name)
	fmt.Println("Field(1):", tData.Field(1).Name)

	fmt.Println("NumMethod:", tData.NumMethod())
	fmt.Println("Method(0):", tData.Method(0).Name)
	fmt.Println("Method(0).NumIn:", tData.Method(0).Type.NumIn())
	fmt.Println("Method(0).NumOut:", tData.Method(0).Type.NumOut())
}
