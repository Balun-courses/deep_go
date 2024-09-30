package main

// go build -gcflags='-m' . |& grep escape

func main() {
	sliceWithStack := make([]byte, 64<<10)
	_ = sliceWithStack

	sliceWithHeap := make([]byte, 64<<10+1)
	_ = sliceWithHeap
}
