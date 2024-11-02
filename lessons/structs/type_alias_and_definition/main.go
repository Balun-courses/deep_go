package main

type (
	A = int
	B int
)

func main() {
	var a A = 1
	var b B = 2

	var ia int = a
	_ = ia

	var ib int = int(b)
	_ = ib
}
