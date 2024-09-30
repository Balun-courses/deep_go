package main

import "fmt"

type Interface interface {
	process(int) bool
}

type String string

func (s String) process(size int) bool {
	return len(s) > size
}

func main() {
	var i String = String("inteface")
	fmt.Println(i.process(10))
	fmt.Println(Interface.process(i, 10))
	fmt.Println(interface{ process(int) bool }.process(i, 10))
}
