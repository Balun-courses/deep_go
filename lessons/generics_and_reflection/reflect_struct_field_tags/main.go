package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	X int  `json:"x" xml:"name"`
	Y bool `json:"y,omitempty"`
}

func main() {
	t := reflect.TypeOf(Data{})
	tXTag := t.Field(0).Tag
	tYTag := t.Field(1).Tag

	fmt.Println(reflect.TypeOf(tXTag), reflect.TypeOf(tYTag))

	value, present := tXTag.Lookup("json")
	fmt.Println(value, present)
	value, present = tYTag.Lookup("json")
	fmt.Println(value, present)
}
