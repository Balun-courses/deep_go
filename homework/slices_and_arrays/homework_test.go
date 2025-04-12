package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	front, rear int
}

// создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	values := make([]int, size)
	return CircularQueue{
		values: values,
		front:  -1,
		rear:   -1,
	}
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.front == -1 {
		q.front = 0
	}

	q.rear = (q.rear + 1) % len(q.values)
	q.values[q.rear] = value

	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		// В очереди только один элемент, сбрасываем её после удаления элемента
		q.front = (q.front + 1) % len(q.values)
	}

	return true
}

// получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.front]
}

// получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.rear]
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.front == -1
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	if (q.front == (q.rear+1)%len(q.values)) || (q.front == 0 && q.rear == len(q.values)-1) {
		return true
	}

	return false
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
