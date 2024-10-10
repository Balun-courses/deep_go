package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

func createPointer() *int {
	value2 := new(int)
	return value2
}

func main() {
	value1 := new(int) // stack
	_ = value1

	value2 := createPointer() // heap
	_ = value2
}
