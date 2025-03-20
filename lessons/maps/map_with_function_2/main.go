package main

import "fmt"

func change(data map[int]int) {
	data = map[int]int{2: 20}
}

func main() {
	data := map[int]int{1: 10}
	fmt.Println(data)
	change(data)
	fmt.Println(data)
}
