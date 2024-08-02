package main

import "fmt"

type Greater interface {
	Hello()
}

type Stranger interface {
	Bye() string
	Greater
	Hello() // overloaded
}

func main() {
	var s Stranger
	fmt.Println(s)
}
