package main

import "fmt"

func main() {
	data := map[string]int{"foo": 0, "bar": 1, "baz": 2}
	for key := range data {
		if key == "foo" {
			delete(data, "bar")
		}
		if key == "bar" {
			delete(data, "foo")
		}
	}

	fmt.Println(data)
}
