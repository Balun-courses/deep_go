package main

// go build -gcflags='-m=3' . |& grep escape

func main() {
	var array [10 * 1024 * 1024]int8
	_ = array
}
