package main

// GOEXPERIMENT=arenas go run main.go
// go run -tags goexperiment.arenas main.go

import (
	"arena"
)

type Data struct {
	value     int
	operaions []int
}

func main() {
	a := arena.NewArena()
	defer a.Free()

	// Arenas will not allocate all reference types automatically
	operations := arena.MakeSlice[int](a, 0, 100)
	data := arena.New[Data](a)
	data.operaions = operations
	_ = data
}
