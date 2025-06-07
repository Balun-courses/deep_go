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
	tasks []Task
	index map[int]int
}

func NewScheduler() Scheduler {
	return Scheduler{
		tasks: []Task{},
		index: make(map[int]int),
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
	taskIndex := len(s.tasks) - 1
	s.heapifyUp(taskIndex)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	index, ok := s.index[taskID]
	if !ok {
		return
	}
	oldPriority := s.tasks[index].Priority
	s.tasks[index].Priority = newPriority

	if newPriority > oldPriority {
		s.heapifyUp(index)
	} else {
		s.heapifyDown(index)
	}
}

func (s *Scheduler) GetTask() Task {
	if len(s.tasks) == 0 {
		return Task{}
	}
	result := s.tasks[0]
	s.tasks[0] = s.tasks[len(s.tasks)-1]
	s.tasks = s.tasks[:len(s.tasks)-1]

	if len(s.tasks) > 0 {
		s.heapifyDown(0)
	}

	return result
}

func (s *Scheduler) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if s.tasks[i].Priority <= s.tasks[parent].Priority {
			break
		}
		s.swap(i, parent)
		i = parent
	}
}

func (s *Scheduler) heapifyDown(i int) {
	heapSize := len(s.tasks)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < heapSize && s.tasks[left].Priority > s.tasks[largest].Priority {
			largest = left
		}
		if right < heapSize && s.tasks[right].Priority > s.tasks[largest].Priority {
			largest = right
		}
		if largest == i {
			break
		}
		s.swap(i, largest)
		i = largest
	}
}

func (s *Scheduler) swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
	s.index[s.tasks[i].Identifier] = i
	s.index[s.tasks[j].Identifier] = j
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
	assert.Equal(t, task1.Identifier, task.Identifier)

	task = scheduler.GetTask()
	assert.Equal(t, task3, task)
}
