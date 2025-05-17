package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	res := strings.Builder{}

	typeOfElem := reflect.TypeOf(person)
	valueOfElem := reflect.ValueOf(&person).Elem()

	n := typeOfElem.NumField()
	for i := range n {
		var omitempty bool
		tag := typeOfElem.Field(i).Tag.Get("properties")
		if tag == "" {
			continue
		}

		split := strings.Split(tag, ",")
		if len(split) > 1 && split[1] == "omitempty" {
			omitempty = true
		}

		value := valueOfElem.Field(i)
		if omitempty && isDefaultValue(value) {
			continue
		}

		res.WriteString(fmt.Sprintf("%v=%v", split[0], value))
		if i != n-1 {
			res.WriteString("\n")
		}
	}

	return res.String()
}

func isDefaultValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int:
		return v.Int() == 0
	case reflect.Bool:
		return !v.Bool()
	default:
		return false
	}
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
