package main

import (
	"fmt"
)

type LazyMap func() map[string]string

func Make(ctr func() map[string]string) LazyMap {
	var initialized bool
	var data map[string]string
	return func() map[string]string {
		if !initialized {
			data = ctr()
			initialized = true
			ctr = nil // for GC
		}

		return data
	}
}

func main() {
	data := Make(func() map[string]string {
		return make(map[string]string)
	})

	fmt.Println(data())
	data()["key"] = "value"
	fmt.Println(data())
}
