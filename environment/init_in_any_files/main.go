package main

import "fmt"

func init() {
	fmt.Println("init[main.go] called")
}

func main() {
	fmt.Println("main called")
	secondary()
}
