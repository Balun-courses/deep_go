package main

import (
	"cmp"
	"fmt"
	"strings"
)

type node[K cmp.Ordered, V any] struct {
	key         K
	value       V
	size        int
	left, right *node[K, V]
}

func size[K cmp.Ordered, V any](current *node[K, V]) int {
	if current == nil {
		return 0
	}

	return current.size
}

func put[K cmp.Ordered, V any](current *node[K, V], k K, v V) *node[K, V] {
	if current == nil {
		return &node[K, V]{key: k, value: v, size: 1}
	}

	if k == current.key {
		current.value = v
		return current
	}

	if k < current.key {
		current.left = put(current.left, k, v)
	} else {
		current.right = put(current.right, k, v)
	}

	current.size = size(current.left) + size(current.right) + 1
	return current
}

func min[K cmp.Ordered, V any](current *node[K, V]) *node[K, V] {
	if current == nil {
		return nil
	}

	if current.left == nil {
		return current
	}

	return min(current.left)
}

func max[K cmp.Ordered, V any](current *node[K, V]) *node[K, V] {
	if current == nil {
		return nil
	}

	if current.right == nil {
		return current
	}

	return max(current.right)
}

// del deletes node from tree
// how it works watch here https://www.youtube.com/watch?v=DkOswl0k7s4
func del[K cmp.Ordered, V any](current *node[K, V], parent *node[K, V], k K) (*node[K, V], bool) {
	if current == nil {
		return nil, false
	}

	var deleted bool

	if k < current.key {
		current.left, deleted = del(current.left, current, k)
		if deleted {
			current.size--
		}

		return current, deleted
	}

	if k > current.key {
		current.right, deleted = del(current.right, current, k)
		if deleted {
			current.size--
		}

		return current, deleted
	}

	// case of leaf
	if current.left == nil && current.right == nil {
		return nil, true
	}

	// case of having one child
	if (current.left != nil && current.right == nil) ||
		(current.right != nil && current.left == nil) {

		// detach child from parent
		var child *node[K, V]
		if current.left != nil {
			child = current.left
			current.left = nil
		} else {
			child = current.right
			current.right = nil
		}

		// "remove" current node from tree, connecting parent whith its grandchild
		if parent != nil {
			if parent.left == current {
				parent.left = child
			} else {
				parent.right = child
			}
		}

		return child, true
	}

	// case of two children when right subtree is bigger
	if current.right.size > current.left.size {
		leaf := min(current.right)
		current.key, leaf.key = leaf.key, current.key
		current.value, leaf.value = leaf.value, current.value

		current.right, _ = del(current.right, current, k)
		current.size--

		return current, true
	}

	// case of two children when left subtree is bigger
	leaf := max(current.left)
	current.key, leaf.key = leaf.key, current.key
	current.value, leaf.value = leaf.value, current.value

	current.left, _ = del(current.left, current, k)
	current.size--

	return current, true
}

func has[K cmp.Ordered, V any](current *node[K, V], k K) bool {
	if current == nil {
		return false
	}

	if k == current.key {
		return true
	}

	if k < current.key {
		return has(current.left, k)
	}

	return has(current.right, k)
}

func inOrderTraverse[K cmp.Ordered, V any](current *node[K, V], fn func(k K, v V)) {
	if current == nil {
		return
	}

	inOrderTraverse(current.left, fn)
	fn(current.key, current.value)
	inOrderTraverse(current.right, fn)
}

func preOrderTraverse[K cmp.Ordered, V any](current *node[K, V], depth int, fn func(k K, v V, d int), fnLeaf func(d int)) {
	if current == nil {
		fnLeaf(depth)
		return
	}

	fn(current.key, current.value, depth)
	preOrderTraverse(current.left, depth+1, fn, fnLeaf)
	preOrderTraverse(current.right, depth+1, fn, fnLeaf)
}

type OrderedMap[K cmp.Ordered, V any] struct {
	root *node[K, V]
}

func NewOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{}
}

func (m *OrderedMap[K, V]) Insert(key K, value V) {
	m.root = put(m.root, key, value)
}

func (m *OrderedMap[K, V]) Erase(key K) {
	m.root, _ = del(m.root, nil, key)
}

func (m *OrderedMap[K, V]) Contains(key K) bool {
	return has(m.root, key)
}

func (m *OrderedMap[K, V]) Size() int {
	return size(m.root)
}

func (m *OrderedMap[K, V]) ForEach(action func(K, V)) {
	inOrderTraverse(m.root, action)
}

func (m *OrderedMap[K, V]) String() string {
	if m.root == nil {
		return "<empty>"
	}

	sb := strings.Builder{}
	preOrderTraverse(m.root, 0, func(k K, v V, d int) {
		sb.WriteString(fmt.Sprintf("%s(%v, %v)\n", strings.Repeat(".", d), k, v))
	}, func(d int) {
		sb.WriteString(fmt.Sprintf("%s<nil>\n", strings.Repeat(".", d)))
	})

	return sb.String()
}
