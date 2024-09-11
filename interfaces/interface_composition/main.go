package main

import "fmt"

type Greater interface {
	Hello()
}

type Stranger interface {
	Hello()
}

func main() {
	var s Stranger
	fmt.Println(s)
}
