package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int

	indexOfLastElement int

	startIndex int
	endIndex   int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:             make([]int, size),
		indexOfLastElement: size - 1,
		startIndex:         -1,
		endIndex:           -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.Empty() {
		q.startIndex = q.incrementIndex(q.startIndex)
	}

	q.endIndex = q.incrementIndex(q.endIndex)
	q.values[q.endIndex] = value

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.startIndex == q.endIndex {
		q.startIndex = -1
		q.endIndex = -1
		return true
	}

	q.startIndex = q.incrementIndex(q.startIndex)

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.startIndex]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.endIndex]
}

func (q *CircularQueue) Empty() bool {
	return q.startIndex == -1 && q.endIndex == -1
}

func (q *CircularQueue) Full() bool {
	if q.Empty() {
		return false
	}

	return q.incrementIndex(q.endIndex) == q.startIndex
}

func (q *CircularQueue) incrementIndex(index int) int {
	index++

	if index > q.indexOfLastElement {
		return 0
	}

	return index
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
