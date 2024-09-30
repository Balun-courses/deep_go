package main

import "fmt"

type account struct {
	balance int
}

type client struct {
	account *account
}

func (c client) add(value int) {
	c.account.balance += value
}

func main() {
	c := client{
		account: &account{},
	}

	c.add(100)
	fmt.Println(c.account.balance)
}
