package main

// GOEXPERIMENT=arenas go run main.go
// go run -tags goexperiment.arenas main.go

import (
	"arena"
)

func main() {
	mem := arena.NewArena()
	defer mem.Free()

	slice := arena.NewSlice[int](mem, 0, 5)
	slice = append(slice, 1, 2, 3, 4, 5)

	slice = append(slice, 6) // moved to heap
}
