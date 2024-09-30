package main

import (
	"fmt"
	"reflect"
)

func main() {
	vData := reflect.ValueOf([]int{1, 2, 3})

	tData := vData.Type()
	t1 := reflect.TypeOf(&[3]int{})
	t2 := reflect.TypeOf([3]int{})

	fmt.Println(tData.ConvertibleTo(t1))
	fmt.Println(tData.ConvertibleTo(t2))

	fmt.Println(vData.CanConvert(t1))
	fmt.Println(vData.CanConvert(t2))
}
