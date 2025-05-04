package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values         []int
	idxLastElement int
	beginIdx       int
	numElem        int
	// need to implement
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:         make([]int, size),
		idxLastElement: size - 1,
		beginIdx:       0,
		numElem:        0,
	} // need to implement
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	firstNotEngagedIndex := q.sumIndex(q.beginIdx, q.numElem)
	q.values[firstNotEngagedIndex] = value

	q.numElem++

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.beginIdx = q.incrementIndex(q.beginIdx)

	q.numElem--

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.beginIdx]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	firstIdxEmpty := q.sumIndex(q.beginIdx, q.numElem)
	return q.values[q.decrementIndex(firstIdxEmpty)]
}

func (q *CircularQueue) Empty() bool {
	return q.numElem == 0
}

func (q *CircularQueue) Full() bool {
	return q.numElem == len(q.values)
}

func (q *CircularQueue) incrementIndex(index int) int {
	index++

	if index > q.idxLastElement {
		return 0
	}

	return index
}

func (q *CircularQueue) decrementIndex(index int) int {
	index--

	if index < 0 {
		return q.idxLastElement
	}

	return index
}

func (q *CircularQueue) sumIndex(index int, valueToAdd int) int {
	index += (valueToAdd % len(q.values))

	if index > q.idxLastElement {
		return index - len(q.values)
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
