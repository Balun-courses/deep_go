package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := new(int)
	*x = 100

	runtime.SetFinalizer(x, func(ptr *int) {
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})

	runtime.SetFinalizer(x, func(ptr *int) { // replace with nil
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})
}
