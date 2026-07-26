package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values   []int
	capacity int
	head     int
	tail     int
	size     int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{make([]int, size), size, 0, 0, 0} // need to implement
}

func (q *CircularQueue) Push(value int) bool {
	if q.size == q.capacity {
		return false
	}

	q.values[q.tail] = value
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.size == 0 {
		return false
	}
	q.head = (q.head + 1) % q.capacity
	q.size--
	return true
}

func (q *CircularQueue) Front() int {
	if q.size == 0 {
		return -1
	}
	return q.values[q.head]
}

func (q *CircularQueue) Back() int {
	if q.size == 0 {
		return -1
	}
	index := (q.tail - 1 + q.capacity) % q.capacity
	return q.values[index]
}

func (q *CircularQueue) Empty() bool {
	return q.size == 0
}

func (q *CircularQueue) Full() bool {
	return q.size == q.capacity

}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
