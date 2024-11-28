package main

import (
	"fmt"
	"reflect"
)

func main() {
	dataStruct := reflect.StructOf([]reflect.StructField{
		{Name: "Age", Type: reflect.TypeOf("abc")},
	})

	fmt.Println("dataStruct:", dataStruct)

	dataArray := reflect.ArrayOf(5, reflect.TypeOf(123))

	fmt.Println("dataArray:", dataArray)

	dataPointer := reflect.PointerTo(dataArray)

	fmt.Println("dataPointer:", dataPointer)
}
