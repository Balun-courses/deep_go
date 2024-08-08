package main

import "fmt"

func main() {
	// incorrect code points
	data := []rune{0x0011FFFF, 0x0012FFFF}
	str := string(data)

	fmt.Println(str)
}
