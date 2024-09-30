package main

import "fmt"

func init() {
	fmt.Println("init[secondary.go] called")
}

func secondary() {
	fmt.Println("secondary called")
}
