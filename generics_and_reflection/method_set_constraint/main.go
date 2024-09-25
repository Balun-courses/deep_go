package main

type Data struct{}

func (d Data) Do1() {}

type Constraint interface {
	Data
	Do2()
}

// generics restriction
func GenericDo[T Constraint](value T) {
	value.Do1() // compilation error
	value.Do2() // ok
}
