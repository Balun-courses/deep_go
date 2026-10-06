package main

import (
	"sync"
)

type union interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T union] struct {
	values []T
	length int
	head   int
	tail   int
	mu     sync.Mutex
	// need to implement
}

func NewCircularQueue[T union](size uint16) *CircularQueue[T] {
	return &CircularQueue[T]{
		values: make([]T, size),
		length: int(size),
		head:   -1,
		tail:   -1,
	} // need to implement
}

func (q *CircularQueue[T]) Push(value T) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isFull() {
		return false
	}

	if q.isEmpty() {
		q.head = 0
	}

	q.tail = (q.tail + 1) % q.length
	q.values[q.tail] = value

	return true // need to implement
}

func (q *CircularQueue[T]) Pop() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isEmpty() {
		return false
	}
	if q.head == q.tail {
		q.reset()
	} else {
		q.head = (q.head + 1) % q.length
	}

	return true // need to implement
}

func (q *CircularQueue[T]) Front() T {
	if q.isEmpty() {
		return -1
	}
	q.mu.Lock()
	result := q.values[q.head]
	q.mu.Unlock()

	return result // need to implement
}

func (q *CircularQueue[T]) Back() T {
	if q.isEmpty() {
		return -1
	}
	q.mu.Lock()
	result := q.values[q.tail]
	q.mu.Unlock()

	return result // need to implement
}

func (q *CircularQueue[T]) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.isEmpty() // need to implement
}

func (q *CircularQueue[T]) Full() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.isFull() // need to implement
}

func (q *CircularQueue[T]) isEmpty() bool {
	return q.head == -1
}

func (q *CircularQueue[T]) isFull() bool {
	return (q.tail+1)%q.length == q.head
}

func (q *CircularQueue[T]) reset() {
	q.head = -1
	q.tail = -1
}
