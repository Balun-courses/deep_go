package main

import (
	"fmt"
	"runtime"
	"time"
	"weak"
)

func main() {
	pointer := new(string)
	weakPtr := weak.Make(pointer)

	runtime.GC()
	time.Sleep(time.Second)

	fmt.Println(weakPtr.Value() == nil)
}
