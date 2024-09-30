package main

type A struct {
	value int
}

func (a A) Print() {}

type B struct {
	A
}

type C struct {
	*B
}

func main() {
	var c C = C{B: &B{A: A{value: 10}}}
	//var c C = C{&B{A{10}}} -> the same

	_ = c.B.A.value
	_ = c.A.value
	_ = c.value

	c.B.A.Print()
	c.B.Print()
	c.Print()
}
