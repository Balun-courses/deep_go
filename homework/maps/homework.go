package main

type node struct {
	key, value  int
	left, right *node
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	var insert func(n **node)
	insert = func(n **node) {
		if *n == nil {
			*n = &node{key: key, value: value}
			m.size++
			return
		}
		if key < (*n).key {
			insert(&(*n).left)
		} else if key > (*n).key {
			insert(&(*n).right)
		} else {
			(*n).value = value // обновляем значение
		}
	}
	insert(&m.root)
}

func (m *OrderedMap) Erase(key int) {
	var erase func(n **node)
	erase = func(n **node) {
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
				// найти наименьший элемент в правом поддереве
				minParent := *n
				min := (*n).right
				for min.left != nil {
					minParent = min
					min = min.left
				}
				(*n).key, (*n).value = min.key, min.value
				// удалить дубликат в правом поддереве
				if minParent.left == min {
					erase(&minParent.left)
				} else {
					erase(&minParent.right)
				}
			}
		}
	}
	erase(&m.root)
}

func (m *OrderedMap) Contains(key int) bool {
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

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	var inorder func(n *node)
	inorder = func(n *node) {
		if n == nil {
			return
		}
		inorder(n.left)
		action(n.key, n.value)
		inorder(n.right)
	}
	inorder(m.root)
}
