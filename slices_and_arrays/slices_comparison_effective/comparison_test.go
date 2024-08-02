package main

import (
	"reflect"
	"testing"
)

// go test -bench=. comparison_test.go

type client struct {
	identifier string
	operations []int
}

func (c *client) equal(other *client) bool {
	if c == other {
		return true
	}

	if c.identifier != other.identifier {
		return false
	}

	if len(c.operations) != len(other.operations) {
		return false
	}

	for i := 0; i < len(c.operations); i++ {
		if c.operations[i] != other.operations[i] {
			return false
		}
	}

	return true
}

func BenchmarkWithEqualFunction(b *testing.B) {
	lhs := client{
		identifier: "xx11yy",
		operations: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	rhs := client{
		identifier: "xx11yy",
		operations: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	for i := 0; i < b.N; i++ {
		_ = lhs.equal(&rhs)
	}
}

func BenchmarkWithReflect(b *testing.B) {
	lhs := client{
		identifier: "xx11yy",
		operations: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	rhs := client{
		identifier: "xx11yy",
		operations: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(lhs, rhs)
	}
}
