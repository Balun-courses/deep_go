package main

import "fmt"

type Data struct {
	value int
}

func Process(i interface{}) {
	fmt.Println("nil:", i == nil)
}

func main() {
	var data *Data = nil
	Process(data)
}
