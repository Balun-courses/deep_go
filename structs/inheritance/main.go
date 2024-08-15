package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Intro() string {
	return p.Name
}

type Woman struct {
	Person
}

func (w *Woman) Intro() string {
	return "Mrs. " + w.Person.Intro()
}

func main() {
	woman := Woman{
		Person: Person{
			Name: "Ekaterina",
		},
	}

	fmt.Println(woman.Intro())
	fmt.Println(woman.Person.Intro())
}
