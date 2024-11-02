package main

import "fmt"

type Data struct {
	value int
}

func (d Data) Value1() int {
	return d.value
}

func (d *Data) Value2() int {
	return d.value
}

func main() {
	data := Data{100}
	pointer := &data

	value1ByData := data.Value1
	value1ByPointer := pointer.Value1

	value2ByData := data.Value2
	value2ByPointer := pointer.Value2

	data.value = 200

	fmt.Println("value1ByData:", value1ByData())
	fmt.Println("value1ByPointer:", value1ByPointer())
	fmt.Println("value2ByData:", value2ByData())
	fmt.Println("value2ByPointer:", value2ByPointer())
}
