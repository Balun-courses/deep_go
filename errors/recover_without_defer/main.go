package main

import "fmt"

func process2() {
	fmt.Println(recover())
	panic("error")
	fmt.Println(recover())
}

func process1() {
	fmt.Println(recover())
	process2()
	fmt.Println(recover())
}

func main() {
	fmt.Println(recover())
	process1()
	fmt.Println(recover())
}
