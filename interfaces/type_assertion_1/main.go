package main

import "fmt"

func main() {
	var value int = 100
	var i interface{} = value

	converted1, ok1 := i.(int)
	if ok1 {
		fmt.Println("converted1 int:", converted1)
	}

	converted2, ok2 := i.(float32)
	if ok2 {
		fmt.Println("converted2 float32:", converted2)
	}

	converted3 := i.(int)
	fmt.Println("converted3 int:", converted3)
	converted4 := i.(float32)
	fmt.Println("converted4 float32:", converted4)
}
