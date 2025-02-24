package main

import "fmt"

// Need to show solution

func double(number int) {
	fmt.Println("address:", &number)
	number *= 2
}

func main() {
	number := 100
	double(number)
	fmt.Println("number: ", number)
	fmt.Println("address:", &number)
}
