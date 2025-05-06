package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	Head  *Node
	Count int
}

type Node struct {
	Key    int
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

func (n *Node) Insert(key, value int) (isNewNode bool) {
	if key == n.Key {
		n.Value = value
		return false
	}

	if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key, Value: value, Parent: n}
			return true
		}
		return n.Left.Insert(key, value)
	} else {
		if n.Right == nil {
			n.Right = &Node{Key: key, Value: value, Parent: n}
			return true
		}
		return n.Right.Insert(key, value)
	}
}

func (n *Node) Find(key int) *Node {
	if n.Key == key {
		return n
	}

	if n.Left != nil && key < n.Key {
		return n.Left.Find(key)
	}

	if n.Right != nil && key > n.Key {
		return n.Right.Find(key)
	}

	return nil
}

func (n *Node) Erase(key int) (deletedNode bool) {
	node := n.Find(key)
	if node == nil {
		return false
	}

	var newNode *Node
	if node.Left == nil && node.Right == nil {
		// конечная нода - просто убираем
		newNode = nil
	} else if node.Left != nil && node.Right == nil {
		// есть лист слева, нет справа - просто подменяем
		newNode = node.Left
	} else if node.Left == nil && node.Right != nil {
		// то же самое как с левым, только с правым
		newNode = node.Right
	} else {
		// есть и правая и левая, тогда берем меньшую из правых
		newNode = node.Right
		for newNode.Left != nil {
			newNode = newNode.Left
		}
		newNode.Parent.Left = nil

		newNode.Left = node.Left
		newNode.Right = node.Right
	}

	if newNode != nil {
		newNode.Parent = node.Parent
	}

	if node.Parent.Left != nil && node.Parent.Left.Key == key {
		node.Parent.Left = newNode
	} else {
		node.Parent.Right = newNode
	}

	return true
}

func (n *Node) ForEach(action func(int, int)) {
	if n.Left != nil {
		n.Left.ForEach(action)
	}

	action(n.Key, n.Value)

	if n.Right != nil {
		n.Right.ForEach(action)
	}
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	if m.Head == nil {
		m.Head = &Node{Key: key, Value: value}
		m.Count = 1
		return
	}
	if m.Head.Insert(key, value) {
		m.Count++
	}
}

func (m *OrderedMap) Erase(key int) {
	if m.Head == nil {
		return
	}
	if m.Head.Erase(key) {
		m.Count--
	}
}

func (m *OrderedMap) Contains(key int) bool {
	if m.Head == nil {
		return false
	}
	return m.Head.Find(key) != nil
}

func (m *OrderedMap) Size() int {
	if m.Head == nil {
		return 0
	}
	return m.Count
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.Head == nil {
		return
	}
	m.Head.ForEach(action)
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
