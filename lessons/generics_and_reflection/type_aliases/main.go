package main

type Set[K comparable] map[K]struct{}
type SetBool Set[bool]
type Comparable comparable
type Constraint interface{ int | uint }

func main() {
	setStr := Set[string]{}
	_ = setStr

	setBool := SetBool{}
	_ = setBool

	// var cmp Comparable -> compilation error
	// var cst Constraint -> compilation error
}
