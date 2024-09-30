package main

import "fmt"

type Data struct {
	Value string
}

func (d Data) GetValue() string {
	return d.Value
}

type ValueGetter interface {
	GetValue() string
}

func Print[T ValueGetter](value T) {
	fmt.Println(value.GetValue())
}
