package main

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() int {
	return int(3.14 * c.radius * c.radius)
}

func main() {
	circle := Circle{}
	var iface interface{} = circle
	shape := iface.(Shape)
	_ = shape
}
