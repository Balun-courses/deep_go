package main

import "fmt"

type Obect struct{}

func (o *Obect) Print() {
	fmt.Println("hello world!!!")
}

func main() {
	var object *Obect
	object.Print()
}
