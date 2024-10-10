package main

// GOEXPERIMENT=arenas go run main.go
// go run -tags goexperiment.arenas main.go

import (
	"arena"
	"fmt"
)

type Data struct {
	deposit int
	credit  int
}

func main() {
	a := arena.NewArena()
	data := arena.New[Data](a)
	a.Free()

	// use after free
	fmt.Println(data)
}
