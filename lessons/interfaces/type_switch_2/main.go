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

var i foobarer = &thing{}

func main() {
	switch v := i.(type) {
	case fooer:
		fmt.Println("fooer:", v)
	case barer:
		fmt.Println("barer:", v)
	case foobarer:
		fmt.Println("foobarer:", v)
	default:
		panic("none of them")
	}
}
