package main

type Interface interface {
	Do()
}

type Object1 struct{}

func (o *Object1) Do() {}

type Object2 struct{}

func (o Object2) Do() {}

func main() {
	var i Interface

	i = &Object1{}

	i = Object2{}
	i = &Object2{}

	_ = i
}
