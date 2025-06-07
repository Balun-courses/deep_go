package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}
type OrderedMap struct {
	size int
	root *Node
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		size: 0,
		root: nil,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	newNode := &Node{value: value, key: key}
	if m.root == nil {
		m.root = newNode
		m.size++
	}
	m.root = m.insertNode(m.root, newNode)
}
func (m *OrderedMap) insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		m.size++
		return newNode
	}
	if newNode.key < root.key {
		root.left = m.insertNode(root.left, newNode)
	} else if newNode.key > root.key {
		root.right = m.insertNode(root.right, newNode)
	} else {
		root.value = newNode.value
	}
	return root
}
func (m *OrderedMap) Erase(key int) {
	if m.Contains(key) {
		m.root = m.eraseNode(m.root, key)
		m.size--
	}
}
func (m *OrderedMap) eraseNode(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = m.eraseNode(root.left, key)
	} else if key > root.key {
		root.right = m.eraseNode(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		minNode := m.findMin(root.right)
		root.key = minNode.key
		root.value = minNode.value
		root.right = m.eraseNode(root.right, minNode.key)
	}
	return root
}
func (m *OrderedMap) findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (m *OrderedMap) Contains(key int) bool {
	return m.containsNode(m.root, key)
}

func (m *OrderedMap) containsNode(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if root.key > key {
		return m.containsNode(root.left, key)
	} else if root.key < key {
		return m.containsNode(root.right, key)
	} else {
		return true
	}
}
func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.forEachNode(m.root, action)
}

func (m *OrderedMap) forEachNode(node *Node, action func(int, int)) {
	if node != nil {
		m.forEachNode(node.left, action)
		action(node.key, node.value)
		m.forEachNode(node.right, action)
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
