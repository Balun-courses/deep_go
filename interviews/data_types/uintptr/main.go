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
	z := new(int)

	ptrX := unsafe.Pointer(x)
	ptrY := unsafe.Pointer(y)
	addressZ := uintptr(unsafe.Pointer(z))

	// arithmetic operation
	_ = addressZ + 2
	_ = addressZ - 2

	runtime.GC()

	*(*int)(ptrX) = 100
	*(*int)(ptrY) = 200
	*(*int)(unsafe.Pointer(addressZ)) = 300 // dangerous
}
