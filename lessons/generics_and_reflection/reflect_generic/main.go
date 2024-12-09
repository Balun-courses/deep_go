package main

import (
	"fmt"
	"reflect"
)

func InvertSlice(args []reflect.Value) []reflect.Value {
	inSlice, length := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, length)
	for idx := length - 1; idx >= 0; idx-- {
		element := inSlice.Index(idx)
		outSlice = reflect.Append(outSlice, element)
	}

	return []reflect.Value{outSlice}
}

func Bind(value interface{}, fn func([]reflect.Value) []reflect.Value) {
	invert := reflect.ValueOf(value).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), fn))
}

func main() {
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{2, 3, 5}))

	var invertStrs func([]string) []string
	Bind(&invertStrs, InvertSlice)
	fmt.Println(invertStrs([]string{"C++", "Go"}))
}
