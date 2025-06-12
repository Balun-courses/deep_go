package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type OrderedMap struct {
	root  *node
	count int
}

type node struct {
	key, value  int
	left, right *node
}

// Erase removes a node with provided key from a tree and returns its root.
func (n *node) Erase(key int) (root *node, erased bool) {
	var prev, cur *node
	for cur = n; cur != nil; {
		if cur.key == key {
			break
		} else if cur.key < key {
			prev = cur
			cur = cur.right
		} else {
			prev = cur
			cur = cur.left
		}
	}

	if cur == nil {
		return n, false // empty tree or node not found
	}

	if cur.right == nil {
		// when right subtree is empty then
		// just set left subtree instead of current node to parent
		if prev == nil {
			return cur.left, true
		} else if prev.key < key {
			prev.right = cur.left
		} else {
			prev.left = cur.left
		}
	} else if cur.left == nil {
		// when left subtree is empty then
		// just set right subtree instead of current node to parent
		if prev == nil {
			return cur.right, true
		} else if prev.key < key {
			prev.right = cur.right
		} else {
			prev.left = cur.right
		}
	} else {
		// search for a node to replace with
		// get the lowest node from right subtree and
		// remove it and set it instead of current node
		lowestKey, lowestValue := cur.right.Lowest()
		cur.right.Erase(*lowestKey)
		if prev == nil {
			return &node{key: *lowestKey, value: *lowestValue, left: cur.left, right: cur.right}, true
		} else if prev.key < key {
			prev.right = &node{key: *lowestKey, value: *lowestValue, left: cur.left, right: cur.right}
		} else {
			prev.left = &node{key: *lowestKey, value: *lowestValue, left: cur.left, right: cur.right}
		}
	}

	return n, true
}

func (n *node) Lowest() (*int, *int) {
	if n == nil {
		return nil, nil
	}

	if n.left == nil {
		return &n.key, &n.value
	} else {
		return n.left.Lowest()
	}
}

func (n *node) Contains(key int) bool {
	if n == nil {
		return false
	}

	if n.key == key {
		return true
	} else if n.key < key {
		return n.right.Contains(key)
	} else {
		return n.left.Contains(key)
	}
}

func (n *node) ForEach(action func(key, value int)) {
	if n == nil {
		return
	}

	n.left.ForEach(action)
	action(n.key, n.value)
	n.right.ForEach(action)
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	var prev *node
	for cur := m.root; cur != nil; {
		if cur.key == key {
			cur.value = value
			return
		} else if cur.key < key {
			prev = cur
			cur = cur.right
		} else {
			prev = cur
			cur = cur.left
		}
	}

	if prev == nil { // the first value
		m.root = &node{key: key, value: value}
	} else if prev.value < value { // new right leaf
		prev.right = &node{key: key, value: value}
	} else { // new left leaf
		prev.left = &node{key: key, value: value}
	}

	m.count += 1
}

func (m *OrderedMap) Erase(key int) {
	root, erased := m.root.Erase(key)
	m.root = root
	if erased {
		m.count -= 1
	}
}

func (m *OrderedMap) Contains(key int) bool {
	return m.root.Contains(key)
}

func (m *OrderedMap) Size() int {
	return m.count
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.root.ForEach(action)
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
