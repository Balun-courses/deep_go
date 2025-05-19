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
	v := reflect.ValueOf(person)
	t := v.Type()

	var properties []string
	fields := make(map[string]string)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("properties")
		if tag == "" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		propName := tagParts[0]
		omitempty := len(tagParts) > 1 && tagParts[1] == "omitempty"

		fieldValue := v.Field(i)
		zero := reflect.Zero(fieldValue.Type()).Interface()
		current := fieldValue.Interface()

		if omitempty && reflect.DeepEqual(current, zero) {
			continue
		}

		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fieldValue.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueStr = fmt.Sprintf("%d", fieldValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			valueStr = fmt.Sprintf("%d", fieldValue.Uint())
		case reflect.Bool:
			valueStr = fmt.Sprintf("%t", fieldValue.Bool())
		default:
			valueStr = fmt.Sprintf("%v", current)
		}

		fields[propName] = valueStr
		properties = append(properties, propName)
	}

	var builder strings.Builder
	for i, prop := range properties {
		if i > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(fmt.Sprintf("%s=%s", prop, fields[prop]))
	}

	return builder.String()
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
