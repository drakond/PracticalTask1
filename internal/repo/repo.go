package repo

import (
	"errors"
	"sync"
	"time"
)

type MemoryStorage struct {
	tasks  map[int]*Task
	nextID int
	mutex  sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (s *MemoryStorage) Create(task *Task) *Task {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task.ID = s.nextID
	s.nextID++

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	s.tasks[task.ID] = task

	return task
}

func (s *MemoryStorage) GetAll() []*Task {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (s *MemoryStorage) GetByID(id int) (*Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func (s *MemoryStorage) Update(id int, task *Task) (*Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existingTask, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status
	existingTask.UpdatedAt = time.Now()

	return existingTask, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}
