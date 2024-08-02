package main

func main() {
	data := make([]int, 2, 6)

	slice1 := data[1:6]
	_ = slice1

	slice2 := data[1:7] // panic
	_ = slice2

	left := 2
	right := 1
	slice3 := data[left:right] // panic
	_ = slice3
}
