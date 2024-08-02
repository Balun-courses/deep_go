package secondary

import "fmt"

func init() {
	fmt.Println("init[secondary.go] called")
}

func Secondary() {
	fmt.Println("secondary called")
}
