package main

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap[C constraints.Ordered, T any] struct {
	head *node[C, T]
	// need to implement
}

type node[C constraints.Ordered, T any] struct {
	left, right *node[C, T]
	key         C
	value       T
	height      int
}

func (n *node[C, T]) rotateLeft() *node[C, T] {
	cur := n.right
	n.right = cur.left
	cur.left = n

	cur.left.updateHeight()
	cur.right.updateHeight()
	cur.updateHeight()

	return cur
}

func (n *node[C, T]) rotateRight() *node[C, T] {
	cur := n.left
	n.left = cur.right
	cur.right = n

	cur.left.updateHeight()
	cur.right.updateHeight()
	cur.updateHeight()

	return cur
}

func (n *node[C, T]) updateHeight() {
	if n == nil {
		return
	}
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}

func (n *node[C, T]) balanceFactor() int {
	return n.right.getHeight() - n.left.getHeight()
}

func (n *node[C, T]) getHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *node[C, T]) rebalance() *node[C, T] {
	if n == nil {
		return n
	}

	n.updateHeight()

	if f := n.balanceFactor(); f == 2 {
		if n.right.balanceFactor() < 0 {
			n.right = n.right.rotateRight()
		}

		n = n.rotateLeft()
	} else if f == -2 {
		if n.left.balanceFactor() < 0 {
			n.left = n.left.rotateLeft()
		}

		n = n.rotateRight()
	}

	return n
}

func (n *node[C, T]) insert(key C, value T) *node[C, T] {
	if n == nil {
		return &node[C, T]{key: key, height: 1}
	}

	if key < n.key {
		n.left = n.left.insert(key, value)
	} else if key > n.key {
		n.right = n.right.insert(key, value)
	} else if key == n.key {
		return n
	}

	return n.rebalance()
}

func (n *node[C, T]) min() *node[C, T] {
	for n.left != nil {
		n = n.left
	}

	return n
}

func (n *node[C, T]) delete(key C) *node[C, T] {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = n.left.delete(key)
	} else if key > n.key {
		n.right = n.right.delete(key)
	} else if n.key == key {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		min := n.right.min()
		n.key, n.value = min.key, min.value
		n.right = n.right.delete(min.key)
	}

	return n.rebalance()
}

func (n *node[C, T]) size() int {
	if n == nil {
		return 0
	}

	left := n.left.size()
	right := n.right.size()

	return left + right + 1
}

func (n *node[C, T]) forEach(action func(C, T)) {
	if n == nil {
		return
	}

	n.left.forEach(action)
	action(n.key, n.value)
	n.right.forEach(action)
}

func NewOrderedMap[C constraints.Ordered, T any]() OrderedMap[C, T] {
	return OrderedMap[C, T]{} // need to implement
}

func (m *OrderedMap[C, T]) Insert(key C, value T) {
	m.head = m.head.insert(key, value)
}

func (m *OrderedMap[C, T]) Erase(key C) {
	m.head = m.head.delete(key)
}

func (m *OrderedMap[C, T]) Contains(key C) bool {
	cur := m.head

	for cur != nil {
		if cur.key == key {
			return true
		} else if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return false
}

func (m *OrderedMap[C, T]) Size() int {
	return m.head.size()
}

func (m *OrderedMap[C, T]) ForEach(action func(C, T)) {
	m.head.forEach(action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[int, int]()
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
