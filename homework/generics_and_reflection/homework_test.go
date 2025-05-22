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
	Name      string   `properties:"name"`
	Address   string   `properties:"address,omitempty"`
	Age       int      `properties:"age"`
	Married   bool     `properties:"married"`
	Phone     *string  `properties:"phone,omitempty"`
	AltPhones []string `properties:"alt_phones,omitempty"`
}

func Serialize(data any) string {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ""
	}
	result := make([]string, 0, v.NumField())
	for i := range v.NumField() {
		tag, ok := v.Type().Field(i).Tag.Lookup("properties")
		if !ok {
			continue
		}
		newTag, omitempty := getPropertyTag(tag)
		if omitempty && v.Field(i).IsZero() {
			continue
		}
		str := strings.Builder{}
		str.WriteString(newTag)
		str.WriteString("=")
		str.WriteString(fieldToString(v.Field(i)))
		result = append(result, str.String())
	}
	return strings.Join(result, "\n")
}

func getPropertyTag(tag string) (string, bool) {
	var omitempty bool
	tagParts := strings.Split(tag, ",")
	newTag := make([]string, 0, len(tagParts))
	for _, tagPart := range tagParts {
		if tagPart == "omitempty" {
			omitempty = true
			continue
		}
		newTag = append(newTag, tagPart)
	}
	return strings.Join(newTag, ","), omitempty
}

func fieldToString(field reflect.Value) string {
	switch field.Kind() {
	case reflect.Pointer:
		return fieldToString(field.Elem())
	case reflect.Slice, reflect.Array:
		result := make([]string, 0, field.Len())
		for i := 0; i < field.Len(); i++ {
			result = append(result, fieldToString(field.Index(i)))
		}
		return strings.Join(result, ",")
	case reflect.String:
		return field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'g', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(field.Bool())
	default:
		panic("unsupported field type")
	}
}

func TestSerialization(t *testing.T) {
	phone := "+7(999)1234567"
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
		"test case with pointer field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
				Phone:   &phone,
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true\nphone=+7(999)1234567",
		},
		"test case with slice field": {
			person: Person{
				Name:      "John Doe",
				Age:       30,
				Married:   true,
				Address:   "Paris",
				AltPhones: []string{"+7(999)7654321", "+7(000)1234567"},
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true\nalt_phones=+7(999)7654321,+7(000)1234567",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
