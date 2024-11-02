package main

type Data struct {
}

func (d Data) Print() {}

func main() {
	var value1 = Data{}
	value1.Print()

	// 1 way
	var value2WithSugar = &Data{}
	value2WithSugar.Print()

	// 2 way
	var value2WithoutSugar = new(Data)
	(*value2WithoutSugar).Print()
}
