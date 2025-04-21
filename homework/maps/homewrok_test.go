package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type TreeNode struct {
	key    int
	value  int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

type OrderedMap struct {
	root *TreeNode
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	newNode := &TreeNode{key: key, value: value}
	if m.root == nil {
		m.root = newNode
		m.size++
		return
	}

	current := m.root
	for {
		if key < current.key {
			if current.left == nil {
				current.left = newNode
				newNode.parent = current
				m.size++
				return
			}
			current = current.left
		} else if key > current.key {
			if current.right == nil {
				current.right = newNode
				newNode.parent = current
				m.size++
				return
			}
			current = current.right
		} else {
			current.value = value
			return
		}
	}
}

func (m *OrderedMap) Erase(key int) {
	node := m.findNode(key)
	if node == nil {
		return
	}

	m.deleteNode(node)
	m.size--
}

func (m *OrderedMap) findNode(key int) *TreeNode {
	current := m.root
	for current != nil {
		if key == current.key {
			return current
		} else if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func (m *OrderedMap) Contains(key int) bool {
	return m.findNode(key) != nil
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.inOrderTraversal(m.root, action)
}

func (m *OrderedMap) deleteNode(node *TreeNode) {
	if node.left == nil && node.right == nil {
		m.replaceNodeInParent(node, nil)
	} else if node.left != nil && node.right != nil {
		changeNode := m.minNode(node.right)
		node.key = changeNode.key
		node.value = changeNode.value
		m.deleteNode(changeNode)
	} else {
		var child *TreeNode
		if node.left != nil {
			child = node.left
		} else {
			child = node.right
		}
		m.replaceNodeInParent(node, child)
	}
}

func (m *OrderedMap) replaceNodeInParent(node, replacement *TreeNode) {
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.left = replacement
		} else {
			node.parent.right = replacement
		}
	} else {
		m.root = replacement
	}

	if replacement != nil {
		replacement.parent = node.parent
	}
}

func (m *OrderedMap) minNode(node *TreeNode) *TreeNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (m *OrderedMap) inOrderTraversal(node *TreeNode, action func(int, int)) {
	if node == nil {
		return
	}
	m.inOrderTraversal(node.left, action)
	action(node.key, node.value)
	m.inOrderTraversal(node.right, action)
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
