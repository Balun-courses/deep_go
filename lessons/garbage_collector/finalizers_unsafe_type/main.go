package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := new(int)
	*x = 100

	runtime.SetFinalizer(x, func(ptr *float64) {
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})
}
