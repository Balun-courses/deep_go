package main

import "fmt"

// GOGC=off go run main.go

var data []byte

func allocate() {
	defer func() {
		fmt.Println(recover())
	}()

	count := 0

	for {
		data = make([]byte, 1<<30)
		for idx := 0; idx < 1<<30; idx += 4096 {
			data[idx] = 100
		}

		fmt.Println("allocated GB:", count)
		count++
	}
}

func main() {
	allocate()
}
