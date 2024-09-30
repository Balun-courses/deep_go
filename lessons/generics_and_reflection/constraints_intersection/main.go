package main

// generics restriction
type _ interface {
	int | ~int // compilation error
}

type _ interface {
	interface{ int } | interface{ ~int } // ok
}

type _ interface {
	int | interface{ ~int } // ok
}
