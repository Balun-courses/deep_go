package main

import "fmt"

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'

func printValue(v interface{}) {
	fmt.Println(v)
	//_, _ = v.(int)
}

func main() {
	var num1 int = 10
	var str1 string = "Hello"

	printValue(num1)
	printValue(str1)

	var num2 int = 10
	var str2 string = "Hello"

	var i interface{}
	i = num2
	i = str2
	_ = i
}
