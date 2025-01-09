package main

type Data struct{}

func MakeData(pointer *int) Data {
	println(*pointer)
	return Data{}
}

func (Data) Print(pointer *int) {
	println(*pointer)
}

func main() {
	var value = 1
	var pointer = &value
	defer MakeData(pointer).Print(pointer)

	value = 2
	pointer = new(int)
	MakeData(pointer)
}
