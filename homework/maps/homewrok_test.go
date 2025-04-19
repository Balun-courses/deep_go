package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	key   int
	value int
	left  *OrderedMap
	right *OrderedMap
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	if m.left == nil && m.right == nil {
		m.key = key
		m.value = value
		m.left = &OrderedMap{}
		m.right = &OrderedMap{}
		return
	}
	if m.key == key {
		return
	}
	if m.key > key {
		m.left.Insert(key, value)
		return
	}
	m.right.Insert(key, value)
}

func (m *OrderedMap) Erase(key int) {
	if m.left == nil && m.right == nil {
		return
	}
	if key < m.key {
		m.left.Erase(key)
		return
	}
	if key > m.key {
		m.right.Erase(key)
		return
	}

	//  If there are no children - we can simply turn the target to it's zero value which is an empty struct
	if m.left.left == nil && m.right.left == nil {
		*m = OrderedMap{}
		return
	}
	// If there is only 1 child - put it in place of the one we're erasing
	if m.right.left == nil {
		*m = *m.left
		return
	}
	if m.left.left == nil {
		*m = *m.right
		return
	}
	//  If there are 2 children we look for the smallest key on the right
	successor := m.right
	for successor.left.left != nil {
		successor = successor.left
	}
	m.key = successor.key
	m.value = successor.value
	successor.Erase(successor.key)
}

func (m *OrderedMap) Contains(key int) bool {
	if m.left == nil && m.right == nil {
		return false
	}
	if m.key == key {
		return true
	}
	if m.key > key {
		return m.left.Contains(key)
	}
	return m.right.Contains(key)
}

func (m *OrderedMap) Size() int {
	if m.right == nil && m.left == nil {
		return 0
	}
	return 1 + m.left.Size() + m.right.Size()
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.left == nil && m.right == nil {
		return
	}
	m.left.ForEach(action)
	action(m.key, m.value)
	m.right.ForEach(action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
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
	//
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
