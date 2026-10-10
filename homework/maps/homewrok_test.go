package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	key, value  int
	left, right *node
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	inserted := false
	m.root, inserted = insertNode(m.root, key, value)
	if inserted {
		m.size++
	}
}

func insertNode(n *node, key, value int) (*node, bool) {
	if n == nil {
		return &node{key: key, value: value}, true
	}

	switch {
	case key < n.key:
		n.left, _ = insertNode(n.left, key, value)
	case key > n.key:
		n.right, _ = insertNode(n.right, key, value)
	default:
		n.value = value
		return n, false
	}

	return n, true
}

func (m *OrderedMap) Erase(key int) {
	removed := false
	m.root, removed = eraseNode(m.root, key)
	if removed {
		m.size--
	}
}

func eraseNode(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}

	removed := false

	switch {
	case key < n.key:
		n.left, removed = eraseNode(n.left, key)
	case key > n.key:
		n.right, removed = eraseNode(n.right, key)
	default:
		removed = true

		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}

		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}

		n.key = successor.key
		n.value = successor.value
		n.right, _ = eraseNode(n.right, successor.key)
	}

	return n, removed
}

func (m *OrderedMap) Contains(key int) bool {
	n := m.root
	for n != nil {
		switch {
		case key < n.key:
			n = n.left
		case key > n.key:
			n = n.right
		default:
			return true
		}
	}

	return false
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	forEach(m.root, action)
}

func forEach(n *node, action func(int, int)) {
	if n == nil {
		return
	}

	forEach(n.left, action)
	action(n.key, n.value)
	forEach(n.right, action)
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
