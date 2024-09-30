package main

import "fmt"

type customer1 struct {
	balance int
}

func (c customer1) add(value int) {
	c.balance += value
}

type customer2 struct {
	balance int
}

func (c *customer2) add(value int) {
	c.balance += value
}

func main() {
	c1 := customer1{}
	c1.add(100)

	c2 := customer2{}
	c2.add(100)

	fmt.Println(c1)
	fmt.Println(c2)
}
