package main

type MyInt int

func (i MyInt) String() string {
	return "number"
}

type Constraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
	any

	// Do()
	// interface{ Do() }
	// ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Do[T Constraint](value T) {
	// ...
}

func main() {
	var value MyInt
	Do(value)
}
