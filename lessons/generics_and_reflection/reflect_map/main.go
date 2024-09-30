package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := map[string]int{"data1": 100, "data2": 200}

	vData := reflect.ValueOf(data)
	vData.SetMapIndex(reflect.ValueOf("data1"), reflect.ValueOf(1000))
	vData.SetMapIndex(reflect.ValueOf("data2"), reflect.ValueOf(2000))

	for it := vData.MapRange(); it.Next(); {
		fmt.Println(it.Key(), it.Value())
	}
}
