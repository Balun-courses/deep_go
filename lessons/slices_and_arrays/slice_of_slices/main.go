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
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func GetData() []Data {
	data := make([]Data, 1000) // ~24KB

	for i := 0; i < len(data); i++ {
		data[i] = Data{
			values: make([]byte, 1<<20),
		}
	}

	clear(data[2:])
	return data[:2]
}

func main() {
	data := GetData()

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}
