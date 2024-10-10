package main

import "fmt"

// go run main.go
// GOGC=off go run main.go

var data []byte

func main() {
	count := 0

	for {
		data = make([]byte, 1<<30)
		/*for idx := 0; idx < 1<<30; idx += 4096 {
			data[idx] = 100
		}*/

		fmt.Println("allocated GB:", count)
		count++
	}
}
