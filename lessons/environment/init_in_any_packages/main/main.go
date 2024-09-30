package main

import (
	"fmt"
	"golang_course/environment/init_in_any_packages/secondary"
)

func init() {
	fmt.Println("init[main.go] called")
}

func main() {
	fmt.Println("main called")
	secondary.Secondary()
}
