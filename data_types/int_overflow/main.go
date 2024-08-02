package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	counter++

	fmt.Println(counter)

	// var counter int32 = math.MaxInt32 + 1 -> compilation error
}
