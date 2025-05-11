package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

type OrderedMap struct {
	root *node
}

type node struct {
	key   int
	value int
	left  *node
	right *node
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = insert(m.root, key, value)
}

func (m *OrderedMap) Erase(key int) {
	if m.root == nil {
		return
	}

	erase(m.root, key)
}

func (m *OrderedMap) Contains(key int) bool {
	if m.root == nil {
		return false
	}

	return contains(m.root, key)
}

func (m *OrderedMap) Size() int {
	size := 0
	inorder(m.root, func(_, _ int) {
		size += 1
	})

	return size
}

func (m *OrderedMap) ForEach(action func(k, v int)) {
	inorder(m.root, action)
}

func contains(n *node, k int) bool {
	if n == nil {
		return false
	}
	if n.key == k {
		return true
	}

	if k < n.key {
		return contains(n.left, k)
	} else {
		return contains(n.right, k)
	}
}

func insert(n *node, k, v int) *node {
	if n == nil {
		return &node{
			key:   k,
			value: v,
			left:  nil,
			right: nil,
		}
	}
	if n.key == k {
		n.value = v
	}

	if k < n.key {
		n.left = insert(n.left, k, v)
	}
	if k > n.key {
		n.right = insert(n.right, k, v)
	}

	return n
}

func inorder(n *node, action func(int, int)) {
	if n == nil {
		return
	}
	inorder(n.left, action)
	action(n.key, n.value)
	inorder(n.right, action)
}

func erase(n *node, k int) *node {
	if n == nil {
		return nil
	}
	if k < n.key {
		n.left = erase(n.left, k)
		return n
	}
	if k > n.key {
		n.right = erase(n.right, k)
		return n
	}

	if n.left == nil {
		return n.right
	}
	if n.right == nil {
		return n.left
	}

	min := n.right
	for min != nil && min.left != nil {
		min = min.left
	}
	n.key = min.key
	n.value = min.value
	n.right = erase(n.right, n.key)

	return n
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
