package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Data struct {
	value string
}

func main() {
	data := Data{
		value: "it's a secret",
	}

	field := reflect.ValueOf(&data).Elem().FieldByName("value")
	pointer := unsafe.Pointer(field.UnsafeAddr())
	strPtr := (*string)(pointer)

	*strPtr = "it's not a secret"
	fmt.Println(data.value)
}
