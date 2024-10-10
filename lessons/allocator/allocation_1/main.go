package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

func getResult() int {
	result := 200
	return result
}

func main() {
	_ = getResult()
}
