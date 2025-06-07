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

func Serialize(person any) string {
	// need to implement
	if person == nil {
		return ""
	}

	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	for v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	str := strings.Builder{}
	for i := range t.NumField() {
		f := t.Field(i)

		tag := f.Tag.Get("properties")
		tags := strings.Split(tag, ",")
		if len(tags) == 0 {
			tags = append(tags, strings.ToLower(f.Name))
		}

		omitempty := false
		if len(tags) > 1 && tags[1] == "omitempty" {
			omitempty = true
		}

		if omitempty && v.Field(i).IsZero() {
			continue
		}

		str.WriteString(fmt.Sprintf("%s=%v\n", tags[0], v.Field(i).Interface()))
	}

	return str.String()[:len(str.String()) - 1]
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
