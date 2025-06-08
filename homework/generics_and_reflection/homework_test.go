package main

import (
	"reflect"
	"strconv"
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

func Serialize(p Person) string {
	var out []string
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)

	for i := 0; i < reflect.TypeOf(p).NumField(); i++ {
		field := v.Field(i)

		tag := t.Field(i).Tag.Get("properties")
		if tag == "" {
			continue
		}

		splitTags := strings.Split(tag, ",")
		if len(splitTags) > 1 && splitTags[1] == "omitempty" && field.IsZero() {
			continue
		}
		if len(splitTags[0]) == 0 {
			splitTags[0] = t.Field(i).Name
		}

		switch field.Kind() {
		case reflect.Bool:
			out = append(out, splitTags[0]+`=`+strconv.FormatBool(field.Bool()))
		case reflect.Float64:
			out = append(out, splitTags[0]+`=`+strconv.FormatFloat(field.Float(), 'f', 0, 64))
		case reflect.String:
			out = append(out, splitTags[0]+`=`+field.String())
		case reflect.Int:
			out = append(out, splitTags[0]+`=`+strconv.FormatInt(field.Int(), 10))
		}
	}

	return strings.Join(out, "\n")
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
