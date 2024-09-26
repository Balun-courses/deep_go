package main

import (
	"fmt"
	"reflect"
)

func main() {
	var z = 100
	var y = &z
	var x interface{} = y

	v := reflect.ValueOf(&x)
	vx := v.Elem()
	vy := vx.Elem()
	vz := vy.Elem()

	fmt.Println(vx, vy, vz)
	vz.Set(reflect.ValueOf(200))
	fmt.Println(vx, vy, vz)
}
