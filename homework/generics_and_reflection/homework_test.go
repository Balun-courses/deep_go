package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	v := reflect.ValueOf(person)
	t := reflect.TypeOf(person)

	var b strings.Builder
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("properties")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		key := parts[0]
		omitempty := len(parts) > 1 && parts[1] == "omitempty"

		// Проверяем на пустое значение при наличии `omitempty`
		if omitempty {
			switch fieldValue.Kind() {
			case reflect.String:
				if fieldValue.String() == "" {
					continue
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if fieldValue.Int() == 0 {
					continue
				}
			case reflect.Bool:
				if !fieldValue.Bool() {
					continue
				}
			}
		}

		// Преобразуем значение в строку
		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fieldValue.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueStr = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Bool:
			valueStr = strconv.FormatBool(fieldValue.Bool())
		default:
			continue
		}

		fmt.Fprintf(&b, "%s=%s\n", key, valueStr)
	}

	// Удаляем последний перевод строки, если есть
	result := b.String()
	return strings.TrimSuffix(result, "\n")
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
