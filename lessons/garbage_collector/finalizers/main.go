package main

import (
	"log"
	"runtime"
	"time"
)

type Data struct {
	name string
}

func MakeFoo(name string) *Data {
	data := &Data{
		name: name,
	}

	runtime.SetFinalizer(data, func(d *Data) {
		log.Println("collected:", d)
	})

	return data
}

func CreateData() {
	data1 := MakeFoo("one")
	data2 := MakeFoo("two")

	log.Println("created:", data1)
	log.Println("created:", data2)
}

func main() {
	for i := 0; i < 3; i++ {
		CreateData()
		time.Sleep(time.Second)
		runtime.GC()
	}
}
