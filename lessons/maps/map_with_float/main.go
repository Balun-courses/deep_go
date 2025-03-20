package main

import "fmt"

func main() {
	var value1 float32 = 0.3
	var value2 float32 = 0.6
	data := make(map[float32]struct{})
	data[value1+value2] = struct{}{}

	var result float32 = 0.9
	_, found := data[result]
	fmt.Println(found)
}
