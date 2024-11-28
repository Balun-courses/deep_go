package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string
	var number int
	var slice []int

	fmt.Println("str:", reflect.ValueOf(str).Comparable())
	fmt.Println("number:", reflect.ValueOf(number).Comparable())
	fmt.Println("slice:", reflect.ValueOf(slice).Comparable())
}
