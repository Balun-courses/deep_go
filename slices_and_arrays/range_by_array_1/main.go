package main

import (
	"fmt"
)

type account struct {
	balance int
}

func main() {
	accounts := [...]account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	// also actual for slices
	for _, a := range accounts {
		a.balance += 50
		//fmt.Println(uintptr(unsafe.Pointer(&a)))
	}

	fmt.Println(accounts)
}
