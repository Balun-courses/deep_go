package main

type Data struct{}

func (d Data) Print()

type StrData interface {
	Data
	String()
}

func main() {
	var data StrData
	_ = data
}
