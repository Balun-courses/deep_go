package main

import (
	"fmt"
)

func process2() {
	defer func() {
		e := recover()
		fmt.Println("#2 recovered:", e)
		panic(e)
	}()

	panic("error")
}

func process1() {
	defer func() {
		fmt.Println("#1 recovered:", recover())
	}()

	process2()
}

func main() {
	process1()
}
