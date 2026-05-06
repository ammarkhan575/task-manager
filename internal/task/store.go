package task

import "fmt"

type Store struct {
	tasks  map[int]*Task
	nextID int
}

type StoreInterface interface {
	NextID() int
	AddTask(t *Task)
	GetAll() []*Task
	GetByID(id int) (*Task, error)
	Delete(id int) bool
}

func NewStore() *Store {
	return &Store{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (s *Store) NextID() int {
	return s.nextID
}

func (s *Store) AddTask(t *Task) {
	t.ID = s.nextID
	s.tasks[t.ID] = t
	s.nextID++
}

func (s *Store) GetAll() []*Task {
	tasks := make([]*Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *Store) GetByID(id int) (*Task, error) {
	if id <= 0 {
		return nil, &ValidationError{Field: "ID", Message: "must be a positive integer"}
	}
	t, ok := s.tasks[id]
	if !ok {
		return nil, fmt.Errorf("GetByID: %d, %w", id, ErrNotFound)
	}
	return t, nil
}

func (s *Store) Delete(id int) bool {
	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}
