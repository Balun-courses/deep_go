package main

// go build -gcflags '-m'

//go:noinline
func getResult() int {
	result := 200
	return result
}

func main() {
	_ = getResult()
}
