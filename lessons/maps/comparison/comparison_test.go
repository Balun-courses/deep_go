package main

import (
	"fmt"
	"reflect"
	"testing"
)

// go test -bench=. comparison_test.go

type client struct {
	operations map[int]int
}

func (c *client) equal(other *client) bool {
	if c == other {
		return true
	}

	if len(c.operations) != len(other.operations) {
		return false
	}

	for key, value := range c.operations {
		otherValue, found := other.operations[key]
		if !found || value != otherValue {
			return false
		}
	}

	return true
}

func BenchmarkWithEqualFunction(b *testing.B) {
	lhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	rhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	for i := 0; i < b.N; i++ {
		_ = lhs.equal(&rhs) // maps.Equal()
	}
}

func BenchmarkWithReflect(b *testing.B) {
	lhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	rhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(lhs, rhs)
	}
}

func BenchmarkWithStringFunction(b *testing.B) {
	lhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	rhs := client{
		operations: map[int]int{1: 2, 2: 5, 3: 7, 4: 9, 5: 10, 6: 3, 7: 4},
	}

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(lhs) == fmt.Sprint(rhs)
	}
}
