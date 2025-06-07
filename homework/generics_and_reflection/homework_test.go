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
	var pairs []string
	var stringResult string

	v := reflect.ValueOf(person)
	t := reflect.TypeOf(person)
	numField := t.NumField()

	for i := 0; i < numField; i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("properties")
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		key := parts[0]
		isOmitempty := len(parts) > 1 && parts[1] == "omitempty"
		if isOmitempty && value.IsZero() {
			continue
		}

		switch value.Kind() {
		case reflect.String:
			stringResult = value.String()
		case reflect.Int:
			stringResult = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			stringResult = fmt.Sprintf("%t", value.Bool())
		default:
			continue
		}
		stringResult = castValue(stringResult)
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, stringResult))
	}

	return strings.Join(pairs, "\n")
}

func castValue(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, "=", `\=`)
	value = strings.ReplaceAll(value, ":", `\:`)
	value = strings.ReplaceAll(value, "\n", `\n`)
	return value
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
