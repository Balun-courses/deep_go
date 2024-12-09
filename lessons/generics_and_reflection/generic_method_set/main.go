package main

type S struct{}

func (S) Bar() {}

type C interface {
	S
	Foo()
}

func foobar[T C](value T) {
	value.Foo()
	value.Bar() // compilation error
}
