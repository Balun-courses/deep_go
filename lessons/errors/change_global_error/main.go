package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("io.EOF == io.EOF:", io.EOF == io.EOF)
	previousErr := io.EOF

	io.EOF = fmt.Errorf("changed error")
	fmt.Println("io.EOF == io.EOF:", io.EOF == io.EOF)
	fmt.Println("previousErr == io.EOF:", previousErr == io.EOF)
}
