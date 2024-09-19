package main

import (
	"fmt"
	"runtime"
)

func panicNil() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic(nil)
}

func panicNilError() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic(new(runtime.PanicNilError))
}

func main() {
	panicNil()
	panicNilError()
}
