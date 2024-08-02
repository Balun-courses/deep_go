package main

import "fmt"

func main() {
	data1 := make([]int, 3, 6)
	data2 := data1[1:3]

	fmt.Println()
	fmt.Println("#1: ", data1, len(data1), cap(data1))
	fmt.Println("#1: ", data2, len(data2), cap(data2))

	data1[1] = 1

	fmt.Println()
	fmt.Println("#2: ", data1, len(data1), cap(data1))
	fmt.Println("#2: ", data2, len(data2), cap(data2))

	data2 = append(data2, 2)

	fmt.Println()
	fmt.Println("#3: ", data1, len(data1), cap(data1))
	fmt.Println("#3: ", data2, len(data2), cap(data2))

	data1 = append(data1, 3)

	fmt.Println()
	fmt.Println("#4: ", data1, len(data1), cap(data1))
	fmt.Println("#4: ", data2, len(data2), cap(data2))

	data2 = append(data2, 4, 5, 6)

	fmt.Println()
	fmt.Println("#5: ", data1, len(data1), cap(data1))
	fmt.Println("#5: ", data2, len(data2), cap(data2))

	data1[1] = 5

	fmt.Println()
	fmt.Println("#6: ", data1, len(data1), cap(data1))
	fmt.Println("#6: ", data2, len(data2), cap(data2))
}
