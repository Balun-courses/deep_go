package main

type Derived1 struct {
	values []int
}

func (d Derived1) Print() {}

type Derived2 struct {
	values []int
}

func (d Derived2) Print() {}

type Base struct {
	Derived1
	Derived2
}

func main() {
	var base Base
	// _ = base.values
	// base.Print()

	base.Derived1.values = nil
	base.Derived2.values = nil

	base.Derived1.Print()
	base.Derived2.Print()
}
