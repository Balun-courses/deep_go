package main

import "fmt"

func main() {
	hello := []byte("Hello ")
	world := "world!"

	helloWorld1 := append(hello, world...) // sugar
	fmt.Println(string(helloWorld1))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)

	copy(helloWorld2[len(hello):], world) // sugar
	fmt.Println(string(helloWorld2))
}
