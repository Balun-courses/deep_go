package main

import (
	"cmp"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node[T cmp.Ordered] struct {
	left, right *Node[T]
	key         T
	value       any
}
type OrderedMap[T cmp.Ordered] struct {
	node *Node[T]
	size int
}

func NewOrderedMap[T cmp.Ordered]() OrderedMap[T] {
	return OrderedMap[T]{}
}

func (m *OrderedMap[T]) Insert(key T, value any) {
	var inserted bool
	m.node, inserted = m.insert(m.node, key, value)
	if inserted {
		m.size++
	}
}

func (m *OrderedMap[T]) insert(node *Node[T], key T, value any) (*Node[T], bool) {
	if node == nil {
		return &Node[T]{key: key, value: value}, true
	}

	var inserted bool
	if key < node.key {
		node.left, inserted = m.insert(node.left, key, value)
	} else if key > node.key {
		node.right, inserted = m.insert(node.right, key, value)
	} else {
		node.value = value
		return node, false
	}
	return node, inserted
}

func (m *OrderedMap[T]) Erase(key T) {
	var deleted bool
	m.node, deleted = m.erase(m.node, key)
	if deleted {
		m.size--
	}
}

func (m *OrderedMap[T]) erase(node *Node[T], key T) (*Node[T], bool) {
	if node == nil {
		return nil, false
	}

	var deleted bool
	if key < node.key {
		node.left, deleted = m.erase(node.left, key)
	} else if key > node.key {
		node.right, deleted = m.erase(node.right, key)
	} else {
		deleted = true
		if node.left == nil {
			return node.right, deleted
		} else if node.right == nil {
			return node.left, deleted
		} else {
			minRight := m.findMin(node.right)
			node.key = minRight.key
			node.value = minRight.value
			node.right, _ = m.erase(node.right, minRight.key)
		}
	}
	return node, deleted
}

func (m *OrderedMap[T]) findMin(node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (m *OrderedMap[T]) Contains(key T) bool {
	node := m.node
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (m *OrderedMap[T]) Size() int {
	return m.size
}

func (m *OrderedMap[T]) ForEach(action func(T, any)) {
	m.inOrder(m.node, action)
}

func (m *OrderedMap[T]) inOrder(node *Node[T], action func(T, any)) {
	if node == nil {
		return
	}
	m.inOrder(node.left, action)
	action(node.key, node.value)
	m.inOrder(node.right, action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[int]()
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
	data.ForEach(func(key int, _ any) {
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
	data.ForEach(func(key int, _ any) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
