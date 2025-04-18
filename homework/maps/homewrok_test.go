package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	left, right *node
	key, value  int
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	var parent *node
	curr := m.root
	for curr != nil {
		parent = curr
		switch {
		case key < curr.key:
			curr = curr.left
		case key > curr.key:
			curr = curr.right
		default: // Key exists, update value
			curr.value = value
			return
		}
	}

	newNode := &node{key: key, value: value}
	if parent == nil { // Tree was empty
		m.root = newNode
	} else if key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	m.size++
}

func findMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func deleteNode(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}

	deleted := false
	if key < n.key {
		n.left, deleted = deleteNode(n.left, key)
	} else if key > n.key {
		n.right, deleted = deleteNode(n.right, key)
	} else {
		deleted = true
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}

		minRight := findMin(n.right)
		n.key = minRight.key
		n.value = minRight.value
		n.right, _ = deleteNode(n.right, minRight.key)
	}
	return n, deleted
}

func (m *OrderedMap) Erase(key int) {
	if newRoot, deleted := deleteNode(m.root, key); deleted {
		m.root = newRoot
		m.size--
	}
}

func (m *OrderedMap) Contains(key int) bool {
	curr := m.root
	for curr != nil {
		switch {
		case key < curr.key:
			curr = curr.left
		case key > curr.key:
			curr = curr.right
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
	traverse(m.root, action)
}

func traverse(n *node, action func(int, int)) {
	if n != nil {
		traverse(n.left, action)
		action(n.key, n.value)
		traverse(n.right, action)
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
