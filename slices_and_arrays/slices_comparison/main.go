package main

import (
	"fmt"
	"reflect"
)

func main() {
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{1, 2, 3, 4, 5}

	// data1 == data2   ->   compilation error

	fmt.Println("equal:", reflect.DeepEqual(data1, data2))
}
