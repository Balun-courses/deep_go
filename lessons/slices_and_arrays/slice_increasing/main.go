package main

func main() {
	slice := make([]int, 0, 3)
	println("slice:", slice, "- slice header address:", &slice)

	slice = append(slice, 1, 2, 3)
	println("slice full capacity:", slice)

	slice = append(slice, 4)
	println("slice after exceed capacity:", slice)
}
