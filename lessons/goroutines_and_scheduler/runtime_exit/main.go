package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("tick")
		}
	}()

	time.Sleep(3 * time.Second)
	runtime.Goexit()
}
