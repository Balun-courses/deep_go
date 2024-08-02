package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data []string
	fmt.Printf("empty=%t nil=%t size=%d\n", len(data) == 0, data == nil, unsafe.Sizeof(data))

	data = []string(nil)
	fmt.Printf("empty=%t nil=%t size=%d\n", len(data) == 0, data == nil, unsafe.Sizeof(data))

	data = []string{}
	fmt.Printf("empty=%t nil=%t size=%d\n", len(data) == 0, data == nil, unsafe.Sizeof(data))

	data = make([]string, 0)
	fmt.Printf("empty=%t nil=%t size=%d\n", len(data) == 0, data == nil, unsafe.Sizeof(data))
}
