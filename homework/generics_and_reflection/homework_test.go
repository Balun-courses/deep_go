package main

import (
	"fmt"
	"reflect"
	"slices"
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
	var res string
	for i := range reflect.TypeOf(person).NumField() {
		field := reflect.TypeOf(person).Field(i)
		tagValues := strings.Split(field.Tag.Get("properties"), ",")
		if len(tagValues) == 0 {
			continue
		}
		name := tagValues[0]
		value := reflect.ValueOf(person).Field(i)

		if slices.Contains(tagValues, "omitempty") && value.String() == "" {
			continue
		}

		if i > 0 {
			res += "\n"
		}
		switch field.Type.Kind() {
		case reflect.Int:
			res += fmt.Sprintf("%s=%d", name, value.Int())
		case reflect.Bool:
			res += fmt.Sprintf("%s=%t", name, value.Bool())
		default:
			res += fmt.Sprintf("%s=%s", name, value)
		}
	}

	return res
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
