package main

import "fmt"

type client struct {
	id         string
	operations []float64
}

func main() {
	client1 := client{id: "x", operations: []float64{1.}}
	client2 := client{id: "y", operations: []float64{1.}}
	fmt.Println(client1 == client2)

	/*
		var anyClient1 any = client{id: "x", operations: []float64{1.}}
		var anyClient2 any = client{id: "y", operations: []float64{1.}}
		fmt.Println(anyClient1 == anyClient2)
	*/
}
