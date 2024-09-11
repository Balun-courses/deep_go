package main

import "fmt"

type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	var a Aboutable = &Book{"Go 101"}
	fmt.Println(a) // &{Go 101}

	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	i = 100
	fmt.Println(i) // 100
}
