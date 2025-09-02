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
	tasks   []*Task
	taskMap map[int]int
}

func NewScheduler() Scheduler {
	return Scheduler{
		tasks:   make([]*Task, 0),
		taskMap: make(map[int]int),
	}
}

func (s *Scheduler) AddTask(task *Task) {
	s.tasks = append(s.tasks, task)
	s.taskMap[task.Identifier] = len(s.tasks) - 1
	s.heapifyUp(len(s.tasks) - 1)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	index, exists := s.taskMap[taskID]
	if !exists {
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

func (s *Scheduler) GetTask() *Task {
	if len(s.tasks) == 0 {
		return nil
	}

	lastIndex := len(s.tasks) - 1
	s.swap(0, lastIndex)

	task := s.tasks[lastIndex]
	delete(s.taskMap, task.Identifier)

	s.tasks = s.tasks[:lastIndex]

	if len(s.tasks) > 0 {
		s.heapifyDown(0)
	}

	return task
}

func (s *Scheduler) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if s.tasks[index].Priority <= s.tasks[parent].Priority {
			break
		}
		s.swap(index, parent)
		index = parent
	}
}

func (s *Scheduler) heapifyDown(index int) {
	lastIndex := len(s.tasks) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		largest := index

		if leftChild <= lastIndex && s.tasks[leftChild].Priority > s.tasks[largest].Priority {
			largest = leftChild
		}

		if rightChild <= lastIndex && s.tasks[rightChild].Priority > s.tasks[largest].Priority {
			largest = rightChild
		}

		if largest == index {
			break
		}

		s.swap(index, largest)
		index = largest
	}
}

func (s *Scheduler) swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
	s.taskMap[s.tasks[i].Identifier] = i
	s.taskMap[s.tasks[j].Identifier] = j
}

func TestTrace(t *testing.T) {
	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler()
	scheduler.AddTask(&task1)
	scheduler.AddTask(&task2)
	scheduler.AddTask(&task3)
	scheduler.AddTask(&task4)
	scheduler.AddTask(&task5)

	task := scheduler.GetTask()
	assert.Equal(t, task5, *task)

	task = scheduler.GetTask()
	assert.Equal(t, task4, *task)

	scheduler.ChangeTaskPriority(1, 100)

	task = scheduler.GetTask()
	assert.Equal(t, task1, *task)

	task = scheduler.GetTask()
	assert.Equal(t, task3, *task)
}
