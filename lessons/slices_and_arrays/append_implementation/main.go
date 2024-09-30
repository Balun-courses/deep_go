package main

import "fmt"

func Append(slice []int, data ...int) []int {
	previousLength := len(slice)
	newLength := previousLength + len(data)

	if newLength > cap(slice) {
		// without smart growth for big slices
		capacity := previousLength * 2
		if capacity == 0 {
			capacity = 1
		}

		newSlice := make([]int, capacity)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:newLength]
	copy(slice[previousLength:newLength], data)
	return slice
}

func main() {
	data := []int{}
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 5)
	data = append(data, 6)
	data = append(data, 7)
	data = append(data, 8)
	data = append(data, 9)

	fmt.Println(data, len(data), cap(data))
}
