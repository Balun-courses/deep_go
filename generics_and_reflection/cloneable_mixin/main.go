package main

import (
	"fmt"
	"reflect"
)

/*
type ParserConstraint interface {
	Parse(string) []string
}

type ParserBase[Derived ParserConstraint] struct {
	impl Derived
}

func (p *ParserBase[Derived]) Parse(data string) []string {
	return p.impl.Parse(data)
}

func Parse[T ParserConstraint](parser ParserBase[T], data string) []string {
	return parser.Parse(data)
}

func main() {
	type JsonParser struct{}
	type XMLParser struct{}

	jsonParser := &ParserBase[JsonParser]{}
	xmlParser := &ParserBase[XMLParser]{}

	Parse(jsonParser, "json data...")
	Parse(xmlParser, "xml data...")
}
*/

type CloneableMixin[T any] struct{}

func (m CloneableMixin[T]) Clone() *T {
	return new(T)
}

type Data1 struct {
	CloneableMixin[Data1]
}

type Data2 struct {
	CloneableMixin[Data2]
}

func main() {
	data1 := Data1{}
	clone1 := data1.Clone()
	fmt.Println(reflect.TypeOf(clone1)) // *main.Data1

	data2 := Data2{}
	clone2 := data2.Clone()
	fmt.Println(reflect.TypeOf(clone2)) // *main.Data2
}
