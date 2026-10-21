package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Node struct{
	key int
	value int
	left *Node
	right *Node
}

func NewNode(key, value int) *Node {
	return &Node{key: key, value: value}
}

type OrderedMap struct {
	root *Node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	link := m.findLink(key)
	if *link != nil {
		(*link).value = value
		return
	}

	*link = NewNode(key, value)
	m.size++
}

func (m *OrderedMap) Erase(key int) {
	link := m.findLink(key)
	if *link == nil {
		return
	}
	node := *link

	if node.left == nil {
		*link = node.right
	} else if node.right == nil {
		*link = node.left
	} else {
		targetLink := &(*link).right
		for (*targetLink).left != nil {
			targetLink = &(*targetLink).left
		}

		node.key = (*targetLink).key
		node.value = (*targetLink).value
		*targetLink = (*targetLink).right
	}
	m.size--
}

func (m *OrderedMap) findLink(key int) **Node {
	link := &m.root

	for *link != nil {
		switch {
		case key > (*link).key:
			link = &(*link).right
		case key < (*link).key:
			link = &(*link).left
		default:
			return link
		}
	}
	return link
}

func (m *OrderedMap) Contains(key int) bool {
	return *m.findLink(key) != nil
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	var walk func(*Node)

	walk = func (node *Node)  {
		if node == nil {
			return
		}
		
		walk(node.left)
		action(node.key, node.value)
		walk(node.right)
	}
	walk(m.root)
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
