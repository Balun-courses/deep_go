package main

import "fmt"

type fooer interface{ foo() }

type thing struct{}

func (t *thing) foo() {}
func (t *thing) bar() {}

func main() {
	var i fooer = &thing{}

	// dynamically checking for certain methods
	_, ok := i.(interface{ bar() })
	fmt.Println("result:", ok)
}
