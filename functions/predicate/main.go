package main

import "fmt"

func Less(lhs, rhs int) bool {
	return lhs < rhs
}

func Greater(lhs, rhs int) bool {
	return lhs > rhs
}

func Sort(numbers []int, predicate func(int, int) bool) {
	for i := 0; i < len(numbers); i++ {
		for j := i; j > 0 && predicate(numbers[j-1], numbers[j]); j-- {
			numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
		}
	}
}

func main() {
	numbers := []int{1, 2, 5, 4, 3}

	Sort(numbers, Less)
	fmt.Println(numbers)

	Sort(numbers, Greater)
	fmt.Println(numbers)
}
