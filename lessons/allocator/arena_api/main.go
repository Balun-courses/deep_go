package main

// GOEXPERIMENT=arenas go run main.go
// go run -tags goexperiment.arenas main.go

import (
	"arena"
)

type Data struct {
	deposit int
	credit  int
}

func main() {
	a := arena.NewArena()
	defer a.Free()

	value := arena.New[int64](a)
	_ = value

	data := arena.New[Data](a)
	_ = data

	slice := arena.MakeSlice[int32](a, 0, 10)
	_ = slice

	cloned := arena.Clone[*Data](data) // moved to heap
	_ = cloned
}
