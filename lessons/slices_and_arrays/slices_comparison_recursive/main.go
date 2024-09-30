package main

import (
	"fmt"
	"reflect"
)

type client struct {
	identifier string
	operations []int
}

func main() {
	data1 := make([]client, 10)
	data2 := make([]client, 10)

	data1[1].operations = append(data1[1].operations, 10)
	fmt.Println("equal:", reflect.DeepEqual(data1, data2))
}
