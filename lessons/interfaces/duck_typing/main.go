package main

import "fmt"

// Need to show solution

type Duck interface {
	Talk() string
	Walk()
	Swim()
}

type Dog struct{}

func (d Dog) Talk() string {
	return "AAAGGGRRR"
}

func (d Dog) Walk() {
	fmt.Println("Walking...")
}

func (d Dog) Swim() {
	fmt.Println("Swimming...")
}

func CatchDuck(duck Duck) {
	fmt.Println("Catching...")
}

func main() {
	dog := Dog{}
	CatchDuck(dog)
}
