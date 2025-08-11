package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type OrderedMap struct {
	root *node
	len  int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) insertNode(root *node, key int, value int) *node {
	if root == nil {
		m.len++
		return &node{key: key, value: value}
	}
	if key < root.key {
		root.left = m.insertNode(root.left, key, value)
	} else if key > root.key {
		root.right = m.insertNode(root.right, key, value)
	} else {
		root.value = value
	}
	return root
}

func (m *OrderedMap) contains(root *node, key int) bool {
	if root == nil {
		return false
	}
	if key == root.key {
		return true
	}
	if key < root.key {
		return m.contains(root.left, key)
	}
	return m.contains(root.right, key)
}

func (m *OrderedMap) removeNode(root *node, key int) *node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = m.removeNode(root.left, key)
	} else if key > root.key {
		root.right = m.removeNode(root.right, key)
	} else {
		m.len--

		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		successor := root.right
		for successor.left != nil {
			successor = successor.left
		}

		root.key = successor.key
		root.value = successor.value

		root.right = m.removeNode(root.right, successor.key)

	}

	return root
}

func (m *OrderedMap) forEach(root *node, action func(int, int)) {
	if root == nil {
		return
	}
	m.forEach(root.left, action)
	action(root.key, root.value)
	m.forEach(root.right, action)
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = m.insertNode(m.root, key, value)
}

func (m *OrderedMap) Erase(key int) {
	m.root = m.removeNode(m.root, key)
}

func (m *OrderedMap) Contains(key int) bool {
	return m.contains(m.root, key)
}

func (m *OrderedMap) Size() int {
	return m.len
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if action != nil {
		m.forEach(m.root, action)
	}
}

func TestOrderedMap(t *testing.T) {
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
