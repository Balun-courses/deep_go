package main

// go build -gcflags='-m' . |& grep escape

func main() {
	sliceWithStack := make([]byte, 64<<10)
	_ = sliceWithStack

	sliceWithHeap := make([]byte, 64<<10+1)
	_ = sliceWithHeap

	// unknown ...
	var sliceSpecialCase = []byte{1 << 30: 1}
	_ = sliceSpecialCase
}
