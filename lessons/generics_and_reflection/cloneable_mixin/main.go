package main

import (
	"fmt"
	"reflect"
)

type CloneableMixin[T any] struct{}

func (m CloneableMixin[T]) Clone() *T {
	return new(T)
}

type Data1 struct {
	CloneableMixin[Data1]
}

type Data2 struct {
	CloneableMixin[Data2]
}

func main() {
	data1 := Data1{}
	clone1 := data1.Clone()
	fmt.Println(reflect.TypeOf(clone1)) // *main.Data1

	data2 := Data2{}
	clone2 := data2.Clone()
	fmt.Println(reflect.TypeOf(clone2)) // *main.Data2
}
