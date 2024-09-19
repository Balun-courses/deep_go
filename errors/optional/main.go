package main

import "fmt"

var NullOptional = Optional[int]{}

type Optional[T any] struct {
	value   T
	present bool
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		value:   value,
		present: true,
	}
}

func (o *Optional[T]) HasValue() bool {
	return o.present
}

func (o *Optional[T]) Value() T {
	return o.value
}

func divide(lhs, rhs int) Optional[int] {
	if rhs == 0 {
		return NullOptional
	}

	result := lhs / rhs
	return NewOptional(result)
}

func main() {
	x := 100
	y := 0

	optional := divide(x, y)
	fmt.Println(optional)
}
