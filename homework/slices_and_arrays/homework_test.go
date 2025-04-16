package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)
	queueGen := NewCircularQueueGen[int](queueSize)

	assert.True(t, queue.Empty())
	assert.True(t, queueGen.Empty())
	assert.False(t, queue.Full())
	assert.False(t, queueGen.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, 0, queueGen.Front())
	assert.Equal(t, -1, queue.Back())
	assert.Equal(t, 0, queueGen.Back())
	assert.False(t, queue.Pop())
	assert.False(t, queueGen.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, queueGen.Push(1))
	assert.True(t, queueGen.Push(2))
	assert.True(t, queueGen.Push(3))
	assert.False(t, queueGen.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queueGen.values))

	assert.False(t, queue.Empty())
	assert.False(t, queueGen.Empty())
	assert.True(t, queue.Full())
	assert.True(t, queueGen.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 1, queueGen.Front())
	assert.Equal(t, 3, queue.Back())
	assert.Equal(t, 3, queueGen.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queueGen.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queueGen.Empty())
	assert.False(t, queue.Full())
	assert.False(t, queueGen.Full())
	assert.True(t, queue.Push(4))
	assert.True(t, queueGen.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))
	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queueGen.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 2, queueGen.Front())
	assert.Equal(t, 4, queue.Back())
	assert.Equal(t, 4, queueGen.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queueGen.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queueGen.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queueGen.Pop())
	assert.False(t, queue.Pop())
	assert.False(t, queueGen.Pop())

	assert.True(t, queue.Empty())
	assert.True(t, queueGen.Empty())
	assert.False(t, queue.Full())
	assert.False(t, queueGen.Full())
}
