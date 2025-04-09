package main

// go build -gcflags='-m' . |& grep escape

func main() {
	sliceWithStack := make([]byte, 64<<10)
	_ = sliceWithStack

	size := 10
	sliceWithHeap1 := make([]byte, size)
	_ = sliceWithHeap1

	sliceWithHeap2 := make([]byte, 64<<10+1)
	_ = sliceWithHeap2

	// unknown ...
	var sliceSpecialCase = []byte{1 << 30: 1}
	_ = sliceSpecialCase
}
