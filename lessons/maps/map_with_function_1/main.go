package main

import "fmt"

func set(data map[int]int, key, value int) {
	data[key] = value
}

func main() {
	data := make(map[int]int)
	fmt.Println(data)
	set(data, 100, 500)
	fmt.Println(data)
}
