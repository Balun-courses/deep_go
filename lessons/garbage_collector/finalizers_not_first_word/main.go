package main

import (
	"fmt"
	"runtime"
	"time"
)

type Foo struct {
	a int
	b string
}

func main() {
	foo := &Foo{a: 1, b: "hello"}
	runtime.SetFinalizer(&foo.b, func(ptr *string) {
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})

	runtime.GC()
	time.Sleep(time.Second)
}
