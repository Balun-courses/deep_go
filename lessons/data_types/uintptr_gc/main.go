package main

import (
	"runtime"
	"unsafe"
)

// go run main.go
// go run -gcflags=-d=checkptr main.go

func main() {
	x := new(int)
	y := new(int)

	ptrX := unsafe.Pointer(x)
	addressY := uintptr(unsafe.Pointer(y))

	// arithmetic operation
	_ = addressY + 2
	_ = addressY - 2

	runtime.GC()

	*(*int)(ptrX) = 100
	*(*int)(unsafe.Pointer(addressY)) = 300 // dangerous
}
