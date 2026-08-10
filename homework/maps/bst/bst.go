package bst

import "cmp"

type Node[T cmp.Ordered, V any] struct {
	key   T
	value V
	left  *Node[T, V]
	right *Node[T, V]
}

func NewNode[T cmp.Ordered, V any](key T, value V, left *Node[T, V], right *Node[T, V]) *Node[T, V] {
	return &Node[T, V]{key, value, left, right}
}

func (n *Node[T, V]) Insert(key T, value V) {
	if key == n.key {
		n.value = value
		return
	}

	child := &n.left
	if key > n.key {
		child = &n.right
	}

	if *child == nil {
		*child = NewNode(key, value, nil, nil)
		return
	}

	(*child).Insert(key, value)
}

func (n *Node[T, V]) Value() V {
	return n.value
}

func (n *Node[T, V]) Key() T {
	return n.key
}

func (n *Node[T, V]) Erase(key T) *Node[T, V] {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = n.left.Erase(key)
		return n
	}

	if key > n.key {
		n.right = n.right.Erase(key)
		return n
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	successor := n.right
	for successor.left != nil {
		successor = successor.left
	}

	n.key = successor.key
	n.value = successor.value
	n.right = n.right.Erase(successor.key)

	return n
}

func (n *Node[T, V]) Contains(key T) bool {
	if n == nil {
		return false
	}

	if key == n.key {
		return true
	}

	if key < n.key {
		return n.left.Contains(key)
	}

	return n.right.Contains(key)
}

func (n *Node[T, V]) ForEach(action func(T, V)) {
	if n == nil {
		return
	}

	n.left.ForEach(action)
	action(n.key, n.value)
	n.right.ForEach(action)
}
