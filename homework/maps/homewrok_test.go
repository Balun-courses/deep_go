package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	root *Node
	size int
}

type Node struct {
	left, right *Node
	key, value  int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

// добавить элемент в словарь
func (m *OrderedMap) Insert(key, value int) {
	var parent *Node
	newNode := Node{
		key:   key,
		value: value,
	}
	if m.root == nil {
		m.root = &newNode
		m.size++
		return
	}

	x := m.root
	for x != nil {
		if x.key == key {
			x = &newNode
			return
		}
		parent = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	if key < parent.key {
		parent.left = &newNode
	} else {
		parent.right = &newNode
	}
	m.size++
}

// удалить элемент из словари
func (m *OrderedMap) Erase(key int) {
	x := m.root
	var parent *Node
	for x != nil {
		if x.key == key {
			break
		}
		parent = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	if x == nil {
		return
	}

	m.size--
	if x.right == nil {
		if parent == nil {
			m.root = x.left
			return
		}
		if x == parent.left {
			parent.left = x.left
			return
		}
		parent.right = x.left
		return
	}

	leftMost := x.right
	parent = nil
	for leftMost.left != nil {
		parent = leftMost
		leftMost = leftMost.left
	}
	if parent != nil {
		parent.left = leftMost.right
	} else {
		x.right = leftMost.right
	}
	x.key = leftMost.key
	x.value = leftMost.value

}

// проверить существование элемента в словаре
func (m *OrderedMap) Contains(key int) bool {
	x := m.root
	for x != nil {
		if x.key == key {
			return true
		}
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	return false
}

// получить количество элементов в словаре
func (m *OrderedMap) Size() int {
	return m.size
}

// применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.root == nil {
		return
	}
	m.forEach(m.root, action)
}

// Центрированный обход (in order, LNR)
// L - левый узел (left), R - правый узел (right), N - родительский узел (node)
func (m *OrderedMap) forEach(n *Node, action func(int, int)) {
	if n.left != nil {
		m.forEach(n.left, action)
	}
	action(n.key, n.value)
	if n.right != nil {
		m.forEach(n.right, action)
	}
}

func TestOrderedMap(t *testing.T) {
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
