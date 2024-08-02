package main

// go build -gcflags='-m=3' . |& grep escape

func main() {
	slice := make([]byte, 64*1024)
	_ = slice
}
