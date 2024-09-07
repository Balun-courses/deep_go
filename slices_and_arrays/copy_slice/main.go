package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println("copied:", dst)
}
