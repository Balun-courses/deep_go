package main

import (
	"fmt"
	"reflect"
)

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(factor int) int {
	return (v.X + v.Y) * factor
}

func main() {
	vector := Vector{X: 5, Y: 15}
	vVector := reflect.ValueOf(vector)

	vAdd := vVector.MethodByName("Add")
	vResults := vAdd.Call([]reflect.Value{reflect.ValueOf(2)})

	fmt.Println(vResults[0].Int())

	negative := func(x int) int {
		return -x
	}

	vNegative := reflect.ValueOf(negative)
	vResults = vNegative.Call([]reflect.Value{reflect.ValueOf(100)})
	fmt.Println(vResults[0].Int())
}
