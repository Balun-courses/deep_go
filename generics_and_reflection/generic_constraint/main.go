package main

// constraint
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Slice[T any] []T

type SliceConstaint[E Unsigned] interface {
	Slice[E]
}

func Do[E Unsigned, T SliceConstaint[E]](values T) {
	// ...
}

func main() {
	var slice1 Slice[uint8]
	var slice2 Slice[uint16]

	Do(slice1)
	Do(slice2)
}
