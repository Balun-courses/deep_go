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
	vf := reflect.ValueOf(person)
	tf := vf.Type()
	numField := vf.NumField()

	res := strings.Builder{}

	for i := 0; i < numField; i++ {
		tag := tf.Field(i).Tag.Get("properties")
		if tag == "" {
			continue
		}
		slicedTag := strings.Split(tag, ",")
		name := slicedTag[0]
		omitEmpty := len(slicedTag) > 1 && slicedTag[1] == "omitempty"

		fieldValue := vf.Field(i)
		if omitEmpty && fieldValue.IsZero() {
			continue
		}

		res.WriteString(name)
		res.WriteString("=")
		res.WriteString(fmt.Sprint(fieldValue.Interface()))
		if i != numField-1 {
			res.WriteString("\n")
		}
	}

	return res.String()
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
