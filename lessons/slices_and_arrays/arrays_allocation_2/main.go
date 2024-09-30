package main

//go:noinline
func allocation() *[10]int {
	var data [10]int
	return &data
}

func main() {
	_ = allocation()
}
