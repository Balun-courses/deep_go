package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "oxo123oxo"
	fmt.Println("TrimRight:", strings.TrimRight(str, "xo"))
	fmt.Println("TrimSuffix:", strings.TrimSuffix(str, "xo"))
	// also for strings.TrimLeft() and strings.TrimPrefix()
}
