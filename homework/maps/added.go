package main

import (
	"runtime"
	"sync"
)

type nodeFin struct {
	key, value  int
	left, right *nodeFin
}

type OrderedMapFin struct {
	root *nodeFin
	size int
	once sync.Once
}

// Создание новой карты и установка финализатора
func NewOrderedMapFin() *OrderedMapFin {
	m := &OrderedMapFin{}
	runtime.SetFinalizer(m, func(m *OrderedMapFin) {
		m.Close()
	})
	return m
}

func (m *OrderedMapFin) Insert(key, value int) {
	var insert func(n **nodeFin)
	insert = func(n **nodeFin) {
		if *n == nil {
			*n = &nodeFin{key: key, value: value}
			m.size++
			return
		}
		if key < (*n).key {
			insert(&(*n).left)
		} else if key > (*n).key {
			insert(&(*n).right)
		} else {
			(*n).value = value
		}
	}
	insert(&m.root)
}

func (m *OrderedMapFin) Erase(key int) {
	var erase func(n **nodeFin)
	erase = func(n **nodeFin) {
		if *n == nil {
			return
		}
		if key < (*n).key {
			erase(&(*n).left)
		} else if key > (*n).key {
			erase(&(*n).right)
		} else {
			m.size--
			if (*n).left == nil {
				*n = (*n).right
			} else if (*n).right == nil {
				*n = (*n).left
			} else {
				// найти минимум в правом поддереве
				parent := *n
				min := (*n).right
				for min.left != nil {
					parent = min
					min = min.left
				}
				(*n).key, (*n).value = min.key, min.value
				if parent.left == min {
					erase(&parent.left)
				} else {
					erase(&parent.right)
				}
			}
		}
	}
	erase(&m.root)
}

func (m *OrderedMapFin) Contains(key int) bool {
	curr := m.root
	for curr != nil {
		if key < curr.key {
			curr = curr.left
		} else if key > curr.key {
			curr = curr.right
		} else {
			return true
		}
	}
	return false
}

func (m *OrderedMapFin) Size() int {
	return m.size
}

func (m *OrderedMapFin) ForEach(action func(int, int)) {
	var inorder func(n *nodeFin)
	inorder = func(n *nodeFin) {
		if n == nil {
			return
		}
		inorder(n.left)
		action(n.key, n.value)
		inorder(n.right)
	}
	inorder(m.root)
}

func (m *OrderedMapFin) clear(n *nodeFin) {
	if n == nil {
		return
	}
	m.clear(n.left)
	m.clear(n.right)
	n.left, n.right = nil, nil
}

// Метод ручного закрытия (и автоматического через финализатор)
func (m *OrderedMapFin) Close() {
	m.once.Do(func() {
		m.clear(m.root)
		m.root = nil
		m.size = 0
	})
}
