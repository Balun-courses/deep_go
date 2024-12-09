package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan string, 2)

	vCh := reflect.ValueOf(ch)
	vCh.Send(reflect.ValueOf("Go"))

	succeeded := vCh.TrySend(reflect.ValueOf("C++"))
	fmt.Println("C++", succeeded)

	succeeded = vCh.TrySend(reflect.ValueOf("Java"))
	fmt.Println("Java", succeeded)

	fmt.Println(vCh.Len(), vCh.Cap())

	value, ok := vCh.Recv()
	fmt.Println(value, ok)

	value, ok = vCh.TryRecv()
	fmt.Println(value, ok)

	value, ok = vCh.TryRecv()
	fmt.Println(value, ok)

	vCh.Close()
}
