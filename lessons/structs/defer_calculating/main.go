package main

import "fmt"

type data1 struct {
	value int
}

func (d data1) print() {
	fmt.Println("data1", d.value)
}

type data2 struct {
	value int
}

func (d *data2) print() {
	fmt.Println("data2", d.value)
}

func main() {
	d1 := data1{}
	defer d1.print()

	d2 := data2{}
	defer d2.print()

	d1.value = 100
	d2.value = 200
}
