package main

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Map(data []int, action func(int) int) []int {
	if data == nil || action == nil {
		return nil
	}

	result := make([]int, len(data))
	for i, value := range data {
		result[i] = action(value)
	}

	return result
}

func Filter(data []int, action func(int) bool) []int {
	if data == nil || action == nil {
		return nil
	}

	result := make([]int, 0, len(data))
	for _, value := range data {
		if action(value) {
			result = append(result, value)
		}
	}

	return result
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	if action == nil {
		return initial
	}

	result := initial
	for _, value := range data {
		result = action(result, value)
	}

	return result
}

func TestMap(t *testing.T) {
	tests := map[string]struct {
		data   []int
		action func(int) int
		result []int
	}{
		"nil numbers": {
			action: func(number int) int {
				return -number
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(number int) int {
				return -number
			},
			result: []int{},
		},
		"inc numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) int {
				return number + 1
			},
			result: []int{2, 3, 4, 5, 6},
		},
		"double numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) int {
				return number * number
			},
			result: []int{1, 4, 9, 16, 25},
		},
		"nil action": {
			data:   []int{1, 2, 3},
			action: nil,
			result: nil,
		},
		"single element": {
			data: []int{42},
			action: func(number int) int {
				return number * 2
			},
			result: []int{84},
		},
		"identity": {
			data: []int{-1, 0, 1},
			action: func(number int) int {
				return number
			},
			result: []int{-1, 0, 1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Map(test.data, test.action)
			assert.True(t, reflect.DeepEqual(test.result, result))
		})
	}
}

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		data   []int
		action func(int) bool
		result []int
	}{
		"nil numbers": {
			action: func(number int) bool {
				return number == 0
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(number int) bool {
				return number == 1
			},
			result: []int{},
		},
		"even numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(number int) bool {
				return number%2 == 0
			},
			result: []int{2, 4},
		},
		"positive numbers": {
			data: []int{-1, -2, 1, 2},
			action: func(number int) bool {
				return number > 0
			},
			result: []int{1, 2},
		},
		"nil action": {
			data:   []int{1, 2, 3},
			action: nil,
			result: nil,
		},
		"nothing matches": {
			data: []int{1, 3, 5},
			action: func(number int) bool {
				return number%2 == 0
			},
			result: []int{},
		},
		"everything matches": {
			data: []int{1, 2, 3},
			action: func(number int) bool {
				return true
			},
			result: []int{1, 2, 3},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Filter(test.data, test.action)
			assert.True(t, reflect.DeepEqual(test.result, result))
		})
	}
}

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		initial int
		data    []int
		action  func(int, int) int
		result  int
	}{
		"nil numbers": {
			action: func(lhs, rhs int) int {
				return 0
			},
		},
		"empty numbers": {
			data: []int{},
			action: func(lhs, rhs int) int {
				return 0
			},
		},
		"sum of numbers": {
			data: []int{1, 2, 3, 4, 5},
			action: func(lhs, rhs int) int {
				return lhs + rhs
			},
			result: 15,
		},
		"sum of numbers with initial value": {
			initial: 10,
			data:    []int{1, 2, 3, 4, 5},
			action: func(lhs, rhs int) int {
				return lhs + rhs
			},
			result: 25,
		},
		"nil action": {
			initial: 7,
			data:    []int{1, 2, 3},
			action:  nil,
			result:  7,
		},
		"product of numbers": {
			initial: 1,
			data:    []int{1, 2, 3, 4},
			action: func(lhs, rhs int) int {
				return lhs * rhs
			},
			result: 24,
		},
		"max of numbers": {
			initial: math.MinInt,
			data:    []int{3, 17, -5, 8},
			action: func(lhs, rhs int) int {
				if rhs > lhs {
					return rhs
				}
				return lhs
			},
			result: 17,
		},
		"order matters": {
			data: []int{1, 2, 3},
			action: func(lhs, rhs int) int {
				return lhs*10 + rhs
			},
			result: 123,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Reduce(test.data, test.initial, test.action)
			assert.Equal(t, test.result, result)
		})
	}
}
