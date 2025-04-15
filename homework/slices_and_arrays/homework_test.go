package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	first  int
	last   int
	size   int
	count  int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		first:  0,
		last:   -1,
		size:   size,
		count:  0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.last = (q.last + 1) % q.size
	q.values[q.last] = value
	q.count++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.first = (q.first + 1) % q.size
	q.count--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.first]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.last]
}

func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

func (q *CircularQueue) Full() bool {
	return q.count == q.size
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
