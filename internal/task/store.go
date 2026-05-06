package task

type Store struct {
	tasks map[int]*Task
	nextID int
}

type StoreInterface interface {
	NextID() int
	AddTask(t *Task)
	GetAll() []*Task
	GetByID(id int) (*Task, bool)
	Delete(id int) bool
}

func NewStore() *Store {
	return &Store{
		tasks: make(map[int]*Task),
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

func (s *Store) GetByID(id int) (*Task, bool) {
	t, ok := s.tasks[id]
	return t, ok
}

func (s *Store) Delete(id int) bool {
	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}