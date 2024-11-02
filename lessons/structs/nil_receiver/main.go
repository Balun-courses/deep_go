package main

import "fmt"

type Obect struct{}

func (o *Obect) Print() {
	if o == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func main() {
	var object *Obect
	object.Print()
}
