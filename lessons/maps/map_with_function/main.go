package main

import "fmt"

func append(data map[int]int, key, value int) {
	data[key] = value
}

func main() {
	data := make(map[int]int)
	fmt.Println(data)
	append(data, 100, 500)
	fmt.Println(data)
}
