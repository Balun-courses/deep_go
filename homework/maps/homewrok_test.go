package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
type OrderedMap struct { ... }

func NewOrderedMap() OrderedMap                      // создать упорядоченный словарь
func (m \*OrderedMap) Insert(key, value int)          // добавить элемент в словарь
func (m \*OrderedMap) Erase(key int)                  // удалить элемент из словари
func (m \*OrderedMap) Contains(key int) bool          // проверить существование элемента в словаре
func (m \*OrderedMap) Size() int                      // получить количество элементов в словаре
func (m \*OrderedMap) ForEach(action func(int, int))  // применить функцию к каждому элементу словаря от меньшего к большему
*/

// go test -v homework_test.go

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = insert(m.root, key, value, &m.size)
}

func insert(n *node, key, value int, size *int) *node {
	if n == nil {
		*size++
		return &node{key: key, value: value}
	}

	switch {
	case key < n.key:
		n.left = insert(n.left, key, value, size)

	case key > n.key:
		n.right = insert(n.right, key, value, size)
	default:
		n.value = value
	}
	return n
}

func (m *OrderedMap) Erase(key int) {
	var del bool
	m.root, del = erase(m.root, key)

	if del {
		m.size--
	}
}

func erase(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}

	var del bool

	switch {
	case key < n.key:
		n.left, del = erase(n.left, key)

	case key > n.key:
		n.right, del = erase(n.right, key)

	default:
		del = true
		if n.left == nil {
			return n.right, true
		}

		if n.right == nil {
			return n.left, true
		}

		temp := n.right
		for temp.left != nil {
			temp = temp.left
		}

		n.key, n.value = temp.key, temp.value
		n.right, _ = erase(n.right, temp.key)

	}

	return n, del

}

func (m *OrderedMap) Contains(key int) bool {
	return contains(m.root, key)
}

func contains(n *node, key int) bool {
	if n == nil {
		return false
	}

	switch {
	case key < n.key:
		return contains(n.left, key)
	case key > n.key:
		return contains(n.right, key)
	case key == n.key:
		return true
	default:
		return false
	}
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	forEach(m.root, action)
}

func forEach(n *node, action func(int, int)) {
	if n == nil {
		return
	}

	forEach(n.left, action)
	action(n.key, n.value)
	forEach(n.right, action)
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
