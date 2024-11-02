package main

import "fmt"

type Data struct {
}

func (d Data) Print() {
	fmt.Println("data")
}

type (
	DataType  Data
	DataAlias = Data
)

func main() {
	data1 := DataType{}
	data1.Print() // compilation error

	data2 := DataAlias{}
	data2.Print()
}
