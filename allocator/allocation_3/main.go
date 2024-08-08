package main

// go build -gcflags '-m'

//go:noinline
func getResult(number *int) int {
	result := *number + 200
	return result
}

func main() {
	number := 100
	_ = getResult(&number)
}
