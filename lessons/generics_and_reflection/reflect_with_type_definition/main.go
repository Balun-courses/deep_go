package main

import (
	"fmt"
	"reflect"
)

type IntAlias = int
type IntDefinition int

func main() {
	var value1 IntAlias = 7
	v1 := reflect.ValueOf(value1)

	fmt.Println("alias type:", v1.Type())
	fmt.Println("alias kind:", v1.Kind())

	var value2 IntDefinition = 7
	v2 := reflect.ValueOf(value2)

	fmt.Println("definition type:", v2.Type())
	fmt.Println("definition kind:", v2.Kind())
}
