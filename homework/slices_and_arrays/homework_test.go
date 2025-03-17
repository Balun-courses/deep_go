package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	maxSize     int
	currentSize int
	beginIndex  int
	endIndex    int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:      make([]int, size),
		maxSize:     size,
		currentSize: 0,
		beginIndex:  0,
		endIndex:    0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.endIndex] = value
	q.currentSize++
	q.incrementEndIndex()
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.incrementBeginIndex()
	q.currentSize--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.beginIndex]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	if q.endIndex == 0 {
		return q.values[q.maxSize-1]
	}

	return q.values[q.endIndex-1]
}

func (q *CircularQueue) Empty() bool {
	return q.currentSize == 0
}

func (q *CircularQueue) Full() bool {
	return q.currentSize == q.maxSize
}

func (q *CircularQueue) incrementEndIndex() {
	if q.endIndex == q.maxSize-1 {
		q.endIndex = 0
		return
	}
	q.endIndex++
}

func (q *CircularQueue) incrementBeginIndex() {
	if q.beginIndex == q.maxSize-1 {
		q.beginIndex = 0
		return
	}

	q.beginIndex++
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
