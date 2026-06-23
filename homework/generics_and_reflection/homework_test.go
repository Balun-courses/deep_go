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

func Serialize(val any) string {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Struct {
		panic("value is not a struct")
	}

	res := make([]string, 0, v.NumField())
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		sf := t.Field(i)
		fv := v.Field(i)

		prop, ok := sf.Tag.Lookup("properties")
		if !ok {
			continue
		}

		propVals := strings.Split(prop, ",")

		omitempty := slices.Contains(propVals[1:], "omitempty")
		if omitempty && fv.IsZero() {
			continue
		}

		res = append(res, propVals[0]+"="+fmt.Sprint(fv.Interface()))
	}

	return strings.Join(res, "\n")
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
