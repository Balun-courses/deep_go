package main

import (
	"fmt"
	"runtime"
)

type Data struct {
	values []byte
}

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func GetData() []Data {
	data := make([]Data, 1000)
	printAllocs()

	for i := 0; i < len(data); i++ {
		data[i] = Data{
			values: make([]byte, 1024*1024),
		}
	}

	printAllocs()
	return data[:2] // need copy or set nil to other elements
}

func main() {
	data := GetData()

	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}
