package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

var globalData *Data

type Data struct {
	name string
}

func NewData(name string) *Data {
	data := &Data{
		name: name,
	}

	fmt.Println("created", name)
	runtime.SetFinalizer(data, func(ptr *Data) {
		fmt.Println("finalizer called on addr", unsafe.Pointer(ptr), "value is", ptr.name)
		globalData = ptr
	})

	return data
}

func main() {
	data := NewData("data")
	_ = data

	runtime.GC()
	time.Sleep(time.Second)

	fmt.Println("resurrected on addr", unsafe.Pointer(globalData), "value is", globalData.name)
}
