package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

// go test -v homework_test.go

type Cards struct {
	Title  string `properties:"title"`
	Number string `properties:"number"`
}

type Person struct {
	Name    string   `properties:"name"`
	Address string   `properties:"address,omitempty"`
	Age     int      `properties:"age"`
	Married bool     `properties:"married"`
	X       []int    `properties:"x,omitempty"`
	Y       [2]int   `properties:"y,omitempty"`
	Z       []string `properties:"z,omitempty"`
	Cards   []Cards  `properties:"cards,omitempty"`
}

func Serialize[T any](v T) string {
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Struct {
		return getFieldValueAsString(reflect.ValueOf(v))
	}

	res := strings.Builder{}

	vValue := reflect.ValueOf(v)

	for i := 0; i < vType.NumField(); i++ {
		name := vType.Field(i).Name
		_ = name
		tag, ok := vType.Field(i).Tag.Lookup("properties")
		if !ok {
			continue
		}

		tagsProperties := strings.Split(tag, ",")
		fieldName := tagsProperties[0]
		omitEmpty := slices.Contains(tagsProperties[1:], "omitempty")

		if vValue.Field(i).IsZero() && omitEmpty {
			continue
		}

		vFieldValue := getFieldValueAsString(reflect.ValueOf(v).Field(i))

		var separator = ""
		if i != 0 {
			separator = "\n"
		}

		res.Write([]byte(separator + fieldName + "=" + vFieldValue))
	}

	return res.String()
}

func getFieldValueAsString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Bool:
		if value.Bool() {
			return "true"
		}
		return "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.Itoa(int(value.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.Itoa(int(value.Uint()))
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', 2, 64)
	case reflect.Array, reflect.Slice:
		var res = strings.Builder{}
		res.Grow(value.Len())
		res.Write([]byte("["))

		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)

			delimiter := ""
			if i != 0 {
				delimiter = ","
			}

			res.Write([]byte(delimiter + getFieldValueAsString(element)))
		}

		res.Write([]byte("]"))

		return res.String()
	case reflect.String:
		return value.String()
	case reflect.Struct:
		return Serialize(value.Interface())
	default:
		return fmt.Sprintf("unsupported type %s", value.Type().String())
	}
}

func TestSerialization(t *testing.T) {
	t.Run("test structs", func(t *testing.T) {
		tests := map[string]struct {
			entity Person
			result string
		}{
			"test case with empty fields": {
				result: "name=\nage=0\nmarried=false",
			},
			"test case with fields": {
				entity: Person{
					Name:    "John Doe",
					Age:     30,
					Married: true,
				},
				result: "name=John Doe\nage=30\nmarried=true",
			},
			"test case with omitempty field": {
				entity: Person{
					Name:    "John Doe",
					Age:     30,
					Married: true,
					Address: "Paris",
				},
				result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
			},
			"test case with empty field and omitempty": {
				entity: Person{
					Name:    "John Doe",
					Age:     30,
					Married: true,
				},
				result: "name=John Doe\nage=30\nmarried=true",
			},
			"test Points": {
				entity: Person{
					X: []int{1, 2, 3},
					Y: [2]int{2, 3},
					Z: []string{"one", "two"},
				},
				result: "name=\nage=0\nmarried=false\nx=[1,2,3]\ny=[2,3]\nz=[one,two]",
			},
			"test Cards": {
				entity: Person{
					Name: "John Doe",
					Age:  30,
					Cards: []Cards{
						{
							Title:  "1",
							Number: "123",
						},
						{
							Title:  "2",
							Number: "234",
						},
					},
				},
				result: "name=John Doe\nage=30\nmarried=false\ncards=[title=1\nnumber=123,title=2\nnumber=234]",
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				result := Serialize(test.entity)
				assert.Equal(t, test.result, result)
			})
		}
	})

	t.Run("test simple values", func(t *testing.T) {
		result := Serialize(123)
		assert.Equal(t, "123", result)

		result = Serialize("123")
		assert.Equal(t, "123", result)

		result = Serialize(true)
		assert.Equal(t, "true", result)

		result = Serialize([]float64{12.3, 54.2})
		assert.Equal(t, "[12.30,54.20]", result)

		result = Serialize(func() {})
		assert.Equal(t, "unsupported type func()", result)
	})
}
