package main

// go build -gcflags='-m' . | grep escape

func main() {
	var arrayWithStack [10 << 20]int8
	_ = arrayWithStack

	var arrayWithHeap [10<<20 + 1]int8
	_ = arrayWithHeap
}
