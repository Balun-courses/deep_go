package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]

	x = append(x, y...)
	x = append(x, y...)

	fmt.Println("a:", a)
	fmt.Println("x:", x)
}
