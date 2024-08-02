package main

type Node[T any] struct {
	Value T
	next  *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

/*
func (n *Node[T]) AddValue[T any](value T) {
	... -> compilation error
}
*/

func main() {
	node1 := Node[int]{}
	node1.Add(&Node[int]{})
	// node1.Add(&Node{Value: 100}) -> compilation error
}
