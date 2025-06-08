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
const (
	TagNameProperties = "properties"
	tagOmitempty      = "omitempty"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	values := reflect.ValueOf(person)
	fieldNum := values.NumField()
	pType := reflect.TypeOf(person)

	dataSlice := make([]string, 0, fieldNum)

	for i := 0; i < fieldNum; i++ {
		field := pType.Field(i)
		tagValue := field.Tag.Get(TagNameProperties)
		tags := strings.Split(tagValue, ",")
		if len(tags) == 0 {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			typeKind := values.Field(i).String()
			if slices.Contains(tags, tagOmitempty) && typeKind == "" {
				continue
			}
			dataSlice = append(dataSlice, fmt.Sprintf("%s=%v", tags[0], typeKind))
		case reflect.Int:
			typeKind := values.Field(i).Int()
			if slices.Contains(tags, tagOmitempty) && typeKind == 0 {
				continue
			}
			dataSlice = append(dataSlice, fmt.Sprintf("%s=%v", tags[0], typeKind))
		case reflect.Bool:
			typeKind := values.Field(i).Bool()
			if slices.Contains(tags, tagOmitempty) && typeKind == false {
				continue
			}
			dataSlice = append(dataSlice, fmt.Sprintf("%s=%v", tags[0], typeKind))
		default:
			panic(fmt.Errorf("unsupported type: %s", field.Type.Kind()))
		}
	}

	return strings.Join(dataSlice, "\n")
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
