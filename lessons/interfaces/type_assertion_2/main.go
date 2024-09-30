package main

import "fmt"

type fooer interface{ foo() }
type barer interface{ bar() }
type foobarer interface {
	foo()
	bar()
}

type thing struct{}

func (t *thing) foo() {}
func (t *thing) bar() {}

func main() {
	var i foobarer = &thing{}
	_, ok := i.(fooer)
	fmt.Println("result:", ok)
}
