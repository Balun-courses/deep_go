package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Serialize(object interface{}) string {
	return "" // need to implement
}

func TestSerializationProperties(t *testing.T) {
	tests := map[string]struct {
		object interface{}
		result string
	}{
		"empty object": {
			object: struct{}{},
			result: "",
		},
		"object without tags": {
			object: struct {
				Count   int
				Title   string
				Pointer *int
			}{},
			result: "",
		},
		"object other tags": {
			object: struct {
				Count   int    `json:"count"`
				Title   string `json:"title"`
				Pointer *int   `json:"pointer"`
			}{},
			result: "",
		},
		"object without values": {
			object: struct {
				Count   int    `properties:"count"`
				Title   string `properties:"title"`
				Pointer *int   `properties:"pointer"`
			}{},
			result: `
				count=
				title=
				pointer=
			`,
		},
		"object without values with omitempty": {
			object: struct {
				Count   int    `properties:"count,omitempty"`
				Title   string `properties:"title,omitempty"`
				Pointer *int   `properties:"pointer,omitempty"`
			}{},
			result: "",
		},
		"object with values": {
			object: struct {
				Count   int    `properties:"count"`
				Title   string `properties:"title"`
				Pointer *int   `properties:"pointer"`
			}{
				Count:   100,
				Title:   "title",
				Pointer: new(int),
			},
			result: `
				count=100
				title=title
				pointer=0
			`,
		},
		"object with other types": {
			object: struct {
				Count   int         `properties:"count"`
				Title   string      `properties:"title"`
				Pointer *int        `properties:"pointer"`
				Slice   []int       `properties:"slice"`
				Map     map[int]int `properties:"map"`
			}{
				Count:   100,
				Title:   "title",
				Pointer: new(int),
				Slice:   []int{1, 2, 3},
				Map:     map[int]int{1: 1},
			},
			result: `
				count=100
				title=title
				pointer=0
			`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.object)
			assert.Equal(t, test.result, result)
		})
	}
}
