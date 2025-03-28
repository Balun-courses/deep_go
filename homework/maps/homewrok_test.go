package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
type Node struct {
	Left  *Node
	Right *Node
	Key   int
	Value int
}

type OrderedMap struct {
	Root *Node
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{} // need to implement
}

func (n *Node) insert(key, value int) {
	if n == nil {
		return
	}

	if n.Key == key {
		return
	}

	if n.Key > key {
		if n.Left != nil {
			n.Left.insert(key, value)
		} else {
			n.Left = &Node{
				Key:   key,
				Value: value,
			}
		}
	} else {
		if n.Right != nil {
			n.Right.insert(key, value)
		} else {
			n.Right = &Node{
				Key:   key,
				Value: value,
			}
		}
	}
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}

	return 1 + n.Left.Size() + n.Right.Size()
}

func (n *Node) contains(key int) bool {
	if n == nil {
		return false
	}

	if n.Key == key {
		return true
	}

	if n.Key > key {
		if n.Left == nil {
			return false
		}

		return n.Left.contains(key)
	}

	if n.Right == nil {
		return false
	}

	return n.Right.contains(key)
}

func (n *Node) forEach(action func(int, int)) {
	if n == nil {
		return
	}

	n.Left.forEach(action)
	action(n.Key, n.Value)
	n.Right.forEach(action)
}

func (n *Node) removeSmallestElem() *Node {
	if n == nil {
		return nil
	}

	minEl := n

	for n.Left != nil {
		minEl = n.Left

		if n.Left.Left == nil {
			n.Left = nil
			break
		}

		n.Left = n.Left.Left
	}

	return minEl
}

func (n *Node) erase(parent *Node, key int) {
	if n == nil {
		return
	}

	if n.Key == key {
		if n.Left == nil && n.Right == nil {
			if parent == n {
				parent = nil
				return
			}

			parent.Left = nil
			parent.Right = nil
			return
		}

		if n.Left != nil && n.Right == nil {
			if parent == n {
				parent = parent.Left
				return
			}

			if parent.Right == n {
				parent.Right = n.Left
			} else {
				parent.Left = n.Left
			}

			return
		}

		if n.Left == nil && n.Right != nil {
			if parent == n {
				parent = parent.Right
				return
			}

			if parent.Right == n {
				parent.Right = n.Right
			} else {
				parent.Left = n.Right
			}

			return
		}
	}

	if n.Right.contains(key) {
		n.Right.erase(n, key)
	} else if n.Left.contains(key) {
		n.Left.erase(n, key)
	}

	return
}

func (m *OrderedMap) Insert(key, value int) {
	if m.Root == nil {
		m.Root = &Node{
			Key:   key,
			Value: value,
		}

		return
	}

	m.Root.insert(key, value)
}

func (m *OrderedMap) Erase(key int) {
	if m == nil || m.Root == nil {
		return
	}

	m.Root.erase(m.Root, key)
}

func (m *OrderedMap) Contains(key int) bool {
	if m == nil || m.Root == nil {
		return false
	}

	if m.Root.Key == key {
		return true
	}

	return m.Root.Left.contains(key) || m.Root.Right.contains(key)
}

func (m *OrderedMap) Size() int {
	if m.Root == nil {
		return 0
	}

	return m.Root.Size()
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m == nil || m.Root == nil {
		return
	}

	m.Root.Left.forEach(action)
	action(m.Root.Key, m.Root.Value)
	m.Root.Right.forEach(action)
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
