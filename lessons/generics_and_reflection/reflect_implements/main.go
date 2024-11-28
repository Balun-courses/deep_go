package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	Method()
}

type Data struct{}

func (c Data) Method() {}

func main() {
	dataType := reflect.TypeOf(Data{})
	interfaceType := reflect.TypeOf((*Interface)(nil)).Elem()
	fmt.Println("implements:", dataType.Implements(interfaceType))
}
