package main

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

const (
	tagName         = "properties"
	suffixOmitEmpty = "omitempty"
)

func Serialize(person Person) string {
	v := reflect.ValueOf(&person).Elem()
	t := v.Type()

	sb := new(strings.Builder)

	for idx := 0; idx < v.NumField(); idx++ {
		field := v.Field(idx)
		tag := t.Field(idx).Tag.Get(tagName)

		if NoTag(tag) || (HasOmitEmpty(tag) && HasZeroValue(field)) {
			continue
		}

		if HasOmitEmpty(tag) {
			tag = strings.TrimSuffix(tag, ","+suffixOmitEmpty)
		}

		sb.WriteString(fmt.Sprintf("%s=%v", tag, field.Interface()))

		if idx < v.NumField()-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()

}

func HasOmitEmpty(tag string) bool {
	return strings.Contains(tag, "omitempty")
}

func HasZeroValue(field reflect.Value) bool {
	return field.IsZero()
}

func NoTag(tag string) bool {
	return tag == ""
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
