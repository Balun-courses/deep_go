package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

func getResult(number *int) int {
	result := *number + 200
	return result
}

func main() {
	number := 100
	_ = getResult(&number)
}
