package main

type A interface {
	Do()
}

type B interface {
	Do()
}

func main() {
	var a A
	var b B

	// Does not work with
	// different packages

	a = b
	b = a
}
