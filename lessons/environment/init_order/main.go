package main

import "fmt"

var value = func() int {
	fmt.Println("value initialized")
	return 0
}()

func init() {
	fmt.Println("init called")
}

func main() {
	fmt.Println("main called")
}
