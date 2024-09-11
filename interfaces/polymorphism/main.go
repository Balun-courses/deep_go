package main

import "fmt"

type Figure interface {
	Area() int
}

type Square struct {
	x int
}

func (s Square) Area() int {
	return s.x * s.x
}

type Rectangle struct {
	x int
	y int
}

func (r Rectangle) Area() int {
	return r.x * r.y
}

func main() {
	figures := []Figure{Square{10}, Rectangle{10, 20}}
	for _, figure := range figures {
		fmt.Println(figure.Area())
	}
}
