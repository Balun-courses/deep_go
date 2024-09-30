package main

import (
	"fmt"
	"math/rand"
	"os"
)

var cache = map[int]struct{}{}

func dirtySum(x, y *int) int {
	result := *x + *y

	// mutations of data (1)
	*x = 0
	*y = 0

	// logging or files/network (2)
	fmt.Println(result)

	// global variables (3)
	cache[result] = struct{}{}

	// dependencies on parameters (4)
	os.Getenv("parameter")

	// unpredictable result (5)
	result += rand.Intn(100)

	return result
}

func pureSum(x, y int) int {
	return x + y
}

func main() {
	x := 13
	y := 12
	fmt.Println(dirtySum(&x, &y))

	fmt.Println(pureSum(13, 12))
}
