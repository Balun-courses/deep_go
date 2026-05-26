package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	front  int
	rear   int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		front:  -1,
		rear:   -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.Empty() {
		q.front = 0
	}

	q.rear = (q.rear + 1) % len(q.values)

	q.values[q.rear] = value

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	// Так зануляем удаленный элемент
	q.values[q.front] = 0

	// если начало и конец указывают на один и тот же элемент, то значит он был последний и надо указатели вернуть в начальное состояние
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		return true
	}

	// если элемент был не последний, то двигаем front вправо по кругу на один шаг
	q.front = (q.front + 1) % len(q.values)

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

func (q *CircularQueue) Empty() bool {
	// -1 это такое специальное состояние означающее что очередь пуста
	return q.front == -1
}

func (q *CircularQueue) Full() bool {
	// если по кругу догнали первый элемент - значит очередь полностью заполнена
	return (q.rear+1)%len(q.values) == q.front
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

	// Проверка корректности Push после полного опустошения очереди
	assert.True(t, queue.Push(1))
	assert.False(t, queue.Empty())
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 1, queue.Back())

	assert.True(t, queue.Push(2))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 2, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 2, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Empty())
}
