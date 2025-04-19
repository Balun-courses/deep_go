package main

import (
	"fmt"
	"runtime"
	"time"
)

type Foo struct {
	bar *Bar
}

type Bar struct {
	foo *Foo
}

func main() {
	foo := &Foo{}
	bar := &Bar{}

	foo.bar = bar
	bar.foo = foo

	runtime.SetFinalizer(foo, func(ptr *Foo) {
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})

	runtime.SetFinalizer(bar, func(ptr *Bar) {
		fmt.Println("finalizer called on addr", ptr, "value is", *ptr)
	})

	runtime.GC()
	time.Sleep(10 * time.Second)
}
