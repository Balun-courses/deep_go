package main

import "fmt"

func main() {
	data := make([]int, 4, 5)
	copy := append(data, 5)

	fmt.Println(data)
	//fmt.Println(data[:5])
	fmt.Println(copy)

	/*
		data := make([]int, 5, 5)
		copy := append(data, 5)

		fmt.Println(data)
		fmt.Println(copy)
	*/
}
