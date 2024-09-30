package main

type Base1[Derived any] struct {
	Derived // compilation error
}

type Base2[Derived any] struct {
	d Derived // ok
}
