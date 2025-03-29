package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v .
func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap[int, int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestOrderedMap2(t *testing.T) {
	t.Run("Empty Map", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		assert.Equal(t, 0, m.Size())
		assert.False(t, m.Contains(1))
	})

	t.Run("Insert and Contains", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(1, "one")
		assert.True(t, m.Contains(1))
		assert.Equal(t, 1, m.Size())
	})

	t.Run("Insert Duplicate Key", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(1, "uno")
		assert.True(t, m.Contains(1))
		assert.Equal(t, 1, m.Size())
	})

	t.Run("Erase Existing Key", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Erase(1)
		assert.False(t, m.Contains(1))
		assert.Equal(t, 0, m.Size())
	})

	t.Run("Erase Non-Existent Key", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Erase(999)
		assert.Equal(t, 0, m.Size())
	})

	t.Run("ForEach Iteration", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(2, "two")
		m.Insert(3, "three")

		values := map[int]string{}
		m.ForEach(func(k int, v string) {
			values[k] = v
		})

		assert.Equal(t, 2, len(values))
		assert.Equal(t, "two", values[2])
		assert.Equal(t, "three", values[3])
	})

	t.Run("Edge Cases", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(0, "zero")
		m.Insert(-1, "negative")
		assert.True(t, m.Contains(0))
		assert.True(t, m.Contains(-1))
	})

	t.Run("Large Input", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		for i := 0; i < 10000; i++ {
			m.Insert(i, "value")
		}
		assert.Equal(t, 10000, m.Size())
	})

	t.Run("Delete Root Node", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(50, "root")
		m.Erase(50)
		assert.False(t, m.Contains(50))
		assert.Equal(t, 0, m.Size())
	})

	t.Run("Delete Leaf Node", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(20000, "leaf")
		assert.True(t, m.Contains(20000))
		m.Erase(20000)
		assert.False(t, m.Contains(20000))
		assert.Equal(t, 0, m.Size())
	})

	t.Run("Delete Node with One Child", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(50000, "parent")
		m.Insert(60000, "child")
		m.Erase(50000)
		assert.False(t, m.Contains(50000))
		assert.True(t, m.Contains(60000))
		assert.Equal(t, 1, m.Size())
	})

	t.Run("Delete Node with Two Children", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		m.Insert(70000, "parent")
		m.Insert(65000, "left child")
		m.Insert(75000, "right child")
		m.Erase(70000)
		assert.False(t, m.Contains(70000))
		assert.True(t, m.Contains(65000))
		assert.True(t, m.Contains(75000))
		assert.Equal(t, 2, m.Size())
	})

	t.Run("Stress Test with Sequential Insert/Delete", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		n := 10000
		for i := 0; i < n; i++ {
			m.Insert(i, "value")
		}
		for i := 0; i < n; i++ {
			m.Erase(i)
		}
		assert.Equal(t, 0, m.Size())
	})

	t.Run("Delete Unexisting", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		assert.False(t, m.Contains(1))
		assert.Equal(t, 0, m.Size())
		m.Erase(1)
		assert.False(t, m.Contains(1))
		assert.Equal(t, 0, m.Size())
	})

	t.Run("Balanced Tree Check", func(t *testing.T) {
		m := NewOrderedMap[int, string]()
		values := []int{50, 25, 75, 10, 30, 60, 90}
		for _, v := range values {
			m.Insert(v, "value")
		}
		for _, v := range values {
			assert.True(t, m.Contains(v))
		}

		var keys []int
		m.ForEach(func(key int, _ string) {
			keys = append(keys, key)
		})
		assert.Equal(t, []int{10, 25, 30, 50, 60, 75, 90}, keys)
	})
}
