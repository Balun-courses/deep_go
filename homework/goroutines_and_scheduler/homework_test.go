package main

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Task struct {
	Identifier int
	Priority   int
	index      int // служебное поле для кучи
}

type taskHeap []*Task

func (h taskHeap) Len() int { return len(h) }

func (h taskHeap) Less(i, j int) bool {
	return h[i].Priority > h[j].Priority // max-heap
}

func (h taskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *taskHeap) Push(x any) {
	task := x.(*Task)
	task.index = len(*h)
	*h = append(*h, task)
}

func (h *taskHeap) Pop() any {
	old := *h
	n := len(old)
	task := old[n-1]
	*h = old[:n-1]
	return task
}

type Scheduler struct {
	heap      taskHeap
	active    map[int]*Task
	originals map[int]Task
}

func NewScheduler() Scheduler {
	return Scheduler{
		heap:      make(taskHeap, 0),
		active:    make(map[int]*Task),
		originals: make(map[int]Task),
	}
}

func (s *Scheduler) AddTask(task Task) {
	if _, exists := s.active[task.Identifier]; exists {
		return
	}
	t := &Task{
		Identifier: task.Identifier,
		Priority:   task.Priority,
	}
	s.originals[task.Identifier] = task
	s.active[task.Identifier] = t
	heap.Push(&s.heap, t)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	if t, ok := s.active[taskID]; ok {
		t.Priority = newPriority
		heap.Fix(&s.heap, t.index)
	}
}

func (s *Scheduler) GetTask() Task {
	if len(s.heap) == 0 {
		return Task{}
	}
	t := heap.Pop(&s.heap).(*Task)
	delete(s.active, t.Identifier)
	original := s.originals[t.Identifier]
	delete(s.originals, t.Identifier)
	return original
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	task := scheduler.GetTask()
	assert.Equal(t, task5, task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, task)

	scheduler.ChangeTaskPriority(1, 100)

	task = scheduler.GetTask()
	assert.Equal(t, task1, task)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}
