package main

import "fmt"

type Square struct{}

func (s *Square) Area() float64 {
	return 0.0
}

func (s *Square) Perimeter() float64 {
	return 0.0
}

const (
	AreaMethod = iota
	PerimeterMethod
)

type Interface struct {
	methods [2]func() float64
}

func NewInterface(square *Square) Interface {
	return Interface{
		methods: [2]func() float64{
			AreaMethod:      square.Area,
			PerimeterMethod: square.Perimeter,
		},
	}
}

func (i *Interface) Area() float64 {
	return i.methods[AreaMethod]()
}

func (i *Interface) Perimeter() float64 {
	return i.methods[PerimeterMethod]()
}

func main() {
	// static dispatch
	square := &Square{}
	fmt.Println(square.Area())
	fmt.Println(square.Perimeter())

	// dynamic dispatch
	iface := NewInterface(square)
	fmt.Println(iface.Area())
	fmt.Println(iface.Perimeter())
}
