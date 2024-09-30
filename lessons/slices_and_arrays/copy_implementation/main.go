package main

import "fmt"

func cp(dst, src []int) {
	minLength := min(len(dst), len(src))
	for idx := 0; idx < minLength; idx++ {
		dst[idx] = src[idx]
	}
}

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)
	copy(dst, src)

	fmt.Println("src", src)
	fmt.Println("dst", dst)
}
