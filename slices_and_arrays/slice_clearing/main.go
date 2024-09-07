package main

import "fmt"

func main() {
	// remove all elements
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first: ", first, " : ", len(first), " : ", cap(first))

	// keep allocated memory
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, " : ", len(second), " : ", cap(second))

	// zero out all elements
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third: ", third, " : ", len(third), " : ", cap(third))

	// zero two elements
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, " : ", len(fourth), " : ", cap(fourth))
}
