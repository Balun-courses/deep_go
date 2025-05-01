package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	insertIndex int
	deleteIndex int
	size        int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:      make([]int, size),
		insertIndex: 0,
		deleteIndex: 0,
		size:        0,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	} else {
		q.size += 1
		q.values[q.insertIndex] = value
		if q.insertIndex == len(q.values)-1 {
			q.insertIndex = 0
		} else {
			q.insertIndex += 1
		}
	}
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	} else {
		if q.values[q.deleteIndex] != 0 {
			q.values[q.deleteIndex] = 0
		}
		if q.deleteIndex == len(q.values)-1 {
			q.deleteIndex = 0
		} else {
			q.deleteIndex += 1
		}
		q.size -= 1
	}
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[(q.deleteIndex)%len(q.values)]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	lastIndex := (q.insertIndex - 1)
	if lastIndex < 0 {
		lastIndex = len(q.values) - 1
	}
	return q.values[lastIndex]
}

func (q *CircularQueue) Empty() bool {
	return q.size == 0
}

func (q *CircularQueue) Full() bool {
	return q.size == len(q.values)
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
