package main

import (
	"cmp"
)

type node[K cmp.Ordered, V any] struct {
	key         K
	value       V
	left, right *node[K, V]
}

type OrderedMap[K cmp.Ordered, V any] struct {
	size int
	root *node[K, V]
}

func (m *OrderedMap[K, V]) put(current *node[K, V], k K, v V) *node[K, V] {
	if current == nil {
		return &node[K, V]{key: k, value: v}
	}

	if k == current.key {
		current.value = v
		return current
	}

	if k < current.key {
		current.left = m.put(current.left, k, v)
	} else {
		current.right = m.put(current.right, k, v)
	}

	return current
}

func (m *OrderedMap[K, V]) del(current *node[K, V], k K) *node[K, V] {
	if current == nil {
		return nil
	}

	if k < current.key {
		current.left = m.del(current.left, k)
		return current
	}

	if k > current.key {
		current.right = m.del(current.right, k)
		return current
	}

	if current.left == nil {
		return current.right
	}

	if current.right == nil {
		return current.left
	}

	// find left with min in right subtree
	leaf := current
	for leaf.left != nil {
		leaf = leaf.left
	}

	current.value, current.key = leaf.value, leaf.key
	current.right = m.del(current.right, current.key)

	return current
}

func (m *OrderedMap[K, V]) has(current *node[K, V], k K) bool {
	if current == nil {
		return false
	}

	if k == current.key {
		return true
	}

	if k < current.key {
		return m.has(current.left, k)
	}

	return m.has(current.right, k)
}

func (m *OrderedMap[K, V]) inOrderTraverse(current *node[K, V], fn func(k K, v V)) {
	if current == nil {
		return
	}

	m.inOrderTraverse(current.left, fn)
	fn(current.key, current.value)
	m.inOrderTraverse(current.right, fn)
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	m.size++
	m.root = m.put(m.root, key, value)
}

// Erase removes the key-value pair with the given key from the map.
// It is a no-op if the key is not present in the map.
func (m *OrderedMap[K, V]) Erase(key K) {
	if !m.has(m.root, key) {
		return
	}
	m.size--
	m.root = m.del(m.root, key)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	return m.has(m.root, key)
}

func (m *OrderedMap[K, V]) Size() int {
	return m.size
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	m.inOrderTraverse(m.root, action)
}
