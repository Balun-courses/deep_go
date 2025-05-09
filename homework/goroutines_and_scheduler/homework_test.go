package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	Tasks   map[int]Task
	MaxHeap []Task
}

func NewScheduler() Scheduler {
	return Scheduler{
		Tasks:   make(map[int]Task),
		MaxHeap: make([]Task, 0),
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.Tasks[task.Identifier] = task
	s.MaxHeap = append(s.MaxHeap, task)

	s.bubbleUp(len(s.MaxHeap) - 1)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	task, exists := s.Tasks[taskID]
	if !exists || task.Priority == newPriority {
		return
	}

	oldPriority := task.Priority
	task.Priority = newPriority
	s.Tasks[taskID] = task

	for k, v := range s.MaxHeap {
		if v.Identifier == taskID {
			s.MaxHeap[k] = s.Tasks[taskID]
			if oldPriority < newPriority {
				s.bubbleUp(k)
			} else {
				s.bubbleDown(k)
			}
		}
	}
}

func (s *Scheduler) bubbleUp(k int) {
	for k > 0 {
		parent := (k - 1) / 2
		if s.MaxHeap[parent].Priority < s.MaxHeap[k].Priority {
			s.MaxHeap[parent], s.MaxHeap[k] = s.MaxHeap[k], s.MaxHeap[parent]
			k = parent
			continue
		}

		break
	}
}

func (s *Scheduler) bubbleDown(k int) {
	for k < len(s.MaxHeap) {
		leftChild := k*2 + 1
		rightChild := k*2 + 2

		if leftChild < len(s.MaxHeap) && s.MaxHeap[leftChild].Priority > s.MaxHeap[k].Priority {
			s.MaxHeap[leftChild], s.MaxHeap[k] = s.MaxHeap[k], s.MaxHeap[leftChild]
			k = leftChild
			continue
		}

		if rightChild < len(s.MaxHeap) && s.MaxHeap[rightChild].Priority > s.MaxHeap[k].Priority {
			s.MaxHeap[rightChild], s.MaxHeap[k] = s.MaxHeap[k], s.MaxHeap[rightChild]
			k = rightChild
			continue
		}

		break
	}
}

func (s *Scheduler) GetTask() Task {
	task := s.MaxHeap[0]
	s.MaxHeap[0], s.MaxHeap[len(s.MaxHeap)-1] = s.MaxHeap[len(s.MaxHeap)-1], s.MaxHeap[0]
	s.MaxHeap = s.MaxHeap[:len(s.MaxHeap)-1]
	delete(s.Tasks, task.Identifier)

	s.bubbleDown(0)

	return task
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

	newPriority := 100
	scheduler.ChangeTaskPriority(1, newPriority)

	task1.Priority = newPriority
	task = scheduler.GetTask()
	assert.Equal(t, task1, task)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}
