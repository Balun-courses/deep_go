package main

type NewType[T any] T // may not be used as the source types

type Set1[K comparable] map[K]struct{}
type Set2[K comparable] = map[K]struct{} // generic type aliases are not supported

type SetBool1 Set1[bool]
type SetBool2 = Set1[bool]

type Comparable1 comparable
type Comparable2 = comparable

type Constraint1 interface{ int | uint }
type Constraint2 = interface{ int | uint }

func main() {
	setStr := Set1[string]{}
	_ = setStr

	setBool := SetBool1{}
	_ = setBool
}
