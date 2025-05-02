package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Node struct {
	left  *Node
	right *Node
	key   int
	value int
}

type OrderedMap struct {
	ptrRoot *Node
}

func NewOrderedMap() OrderedMap {
	o := OrderedMap{}
	return o
}

func (m *OrderedMap) Insert(key, value int) {
	m.ptrRoot = m.insert(m.ptrRoot, key, value)
}
func (m *OrderedMap) insert(node *Node, key, value int) *Node {
	if node == nil {
		return &Node{key: key, value: value}
	}

	if node.key == key {
		node.value = value
	} else if node.key > key {
		node.left = m.insert(node.left, key, value)
	} else if node.key < key {
		node.right = m.insert(node.right, key, value)
	}
	return node
}
func (m *OrderedMap) Erase(key int) {
	if m.ptrRoot.key == key {
		if m.ptrRoot.left != nil {
			*m.ptrRoot = *m.ptrRoot.left
		} else if m.ptrRoot.right != nil {
			*m.ptrRoot = *m.ptrRoot.right
		}
		return
	}
	m.erase(m.ptrRoot, key)
}
func (m *OrderedMap) erase(node *Node, key int) {
	if node == nil {
		return
	}

	if node.left != nil && node.left.key == key {
		var l *Node
		var r *Node
		if node.left.left != nil {
			l = node.left.left
		}
		if node.left.right != nil {
			r = node.left.right
		}
		if l != nil && r != nil {
			node.left = l
			node.right = r
		} else if l != nil {
			node.left = l
		} else if r != nil {
			node.left = r
		} else {
			node.left = nil
		}
		return
	}
	if node.right != nil && node.right.key == key {
		var l *Node
		var r *Node
		if node.right.left != nil {
			l = node.right.left
		}
		if node.right.right != nil {
			r = node.right.right
		}
		if l != nil && r != nil {
			node.right = l
			node.left = r
		} else if l != nil {
			node.right = l
		} else if r != nil {
			node.right = r
		} else {
			node.right = nil
		}
		return
	}
	if node.key > key {
		m.erase(node.left, key)
	} else {
		m.erase(node.right, key)
	}
}

func (m *OrderedMap) Contains(key int) bool {
	return m.contains(m.ptrRoot, key)
}
func (m *OrderedMap) contains(node *Node, key int) bool {
	if node.key == key {
		return true
	}

	if node.key <= key {
		if node.right == nil {
			return false
		}
		return m.contains(node.right, key)
	} else {
		if node.left == nil {
			return false
		}
		return m.contains(node.left, key)
	}
}

func (m *OrderedMap) Size() int {
	a := 0
	if m.ptrRoot != nil {
		a = 1
	}
	return m.size(m.ptrRoot) + a
}
func (m *OrderedMap) size(node *Node) int {
	if node == nil {
		return 0
	}
	left := 0
	right := 0
	if node.left != nil {
		left = 1
		left += m.size(node.left)
	}
	if node.right != nil {
		right = 1
		right += m.size(node.right)
	}
	return left + right
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.traversal(m.ptrRoot, action)
}
func (m *OrderedMap) traversal(node *Node, action func(int, int)) {
	if node == nil {
		return
	}
	m.traversal(node.left, action)
	action(node.key, node.value)
	m.traversal(node.right, action)
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
