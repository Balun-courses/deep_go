package main

func filterEvenNumbers(numbers []int) []int {
	return nil // implemented
}

func main() {
	// imperative way
	var imperative []int
	for _, number := range []int{1, 2, 3, 4, 5} {
		if number%2 != 0 {
			imperative = append(imperative, number)
		}
	}

	// declarative way
	declarative := filterEvenNumbers([]int{1, 2, 3, 4, 5})
	_ = declarative
}
