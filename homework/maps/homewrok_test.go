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
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = m.insertRecursive(m.root, key, value)
}

func (m *OrderedMap) insertRecursive(n *node, key, value int) *node {
	if n == nil {
		m.size++
		return &node{key: key, value: value}
	}

	if key < n.key {
		n.left = m.insertRecursive(n.left, key, value)
	} else if key > n.key {
		n.right = m.insertRecursive(n.right, key, value)
	} else {
		n.value = value
	}
	return n
}

func (m *OrderedMap) Erase(key int) {
	m.root = m.eraseRecursive(m.root, key)
}

func (m *OrderedMap) eraseRecursive(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = m.eraseRecursive(n.left, key)
	} else if key > n.key {
		n.right = m.eraseRecursive(n.right, key)
	} else {
		m.size--

		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minNode := m.findMin(n.right)
		n.key = minNode.key
		n.value = minNode.value
		n.right = m.eraseRecursive(n.right, minNode.key)
	}
	return n
}

func (m *OrderedMap) findMin(n *node) *node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (m *OrderedMap) Contains(key int) bool {
	return m.containsRecursive(m.root, key)
}

func (m *OrderedMap) containsRecursive(n *node, key int) bool {
	if n == nil {
		return false
	}

	if key < n.key {
		return m.containsRecursive(n.left, key)
	} else if key > n.key {
		return m.containsRecursive(n.right, key)
	}
	return true
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.inOrderTraversal(m.root, action)
}

func (m *OrderedMap) inOrderTraversal(n *node, action func(int, int)) {
	if n != nil {
		m.inOrderTraversal(n.left, action)
		action(n.key, n.value)
		m.inOrderTraversal(n.right, action)
	}
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
