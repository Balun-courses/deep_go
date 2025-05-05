package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"reflect"
)

// go test -v homework_test.go

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type OrderedMap struct {
	root *Node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	if m.root == nil {
		m.root = &Node{key: key, value: value}
		m.size++
		return
	}
	_, added := insert(m.root, key, value)
	if added {
		m.size++
	}
}

func (m *OrderedMap) Erase(key int) {
	if m.root == nil {
		return
	}
	node := erase(m.root, key)
	if node != nil {
		m.size--
	}

}

func (m *OrderedMap) Contains(key int) bool {
	if m.root == nil {
		return false
	}
	return find(m.root, key) != nil
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.root == nil {
		return
	}
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		traverse(node.left)
		action(node.key, node.value)
		traverse(node.right)
	}
	traverse(m.root)
}

func insert(node *Node, key, value int) (*Node, bool) {
	if node == nil {
		return &Node{key: key, value: value}, true
	}
	if node.key == key {
		node.value = value
		return node, false
	}
	var added bool
	if key < node.key {
		node.left, added = insert(node.left, key, value)
	} else {
		node.right, added = insert(node.right, key, value)
	}

	return node, added
}

func find(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	}
	if key < node.key {
		return find(node.left, key)
	}
	return find(node.right, key)
}

func erase(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = erase(node.left, key)
		return node
	}
	if key > node.key {
		node.right = erase(node.right, key)
		return node
	}

	if node.left == nil {
		return node.right
	}

	if node.right == nil {
		return node.left
	}

	minNode := node.right
	for minNode.left != nil {
		minNode = minNode.left
	}

	node.key, node.value = minNode.key, minNode.value
	node.right = erase(node.right, minNode.key)

	return node
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
