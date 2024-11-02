package main

import "fmt"

type Data struct{}

func (d Data) Print() {
	fmt.Println("data")
}

func main() {
	var data Data

	data.Print()
	(&data).Print()

	(Data).Print(data)
	(*Data).Print(&data)
}
