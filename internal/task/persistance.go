package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultFile = "tasks.json"

func (s *Store) Save(path string) error {
	tasks := s.GetAll()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("save: marshal: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil { // 0644 = owner rw, others r.
		return fmt.Errorf("save: write file: %s %w", path, err)
	}

	return nil
}

func (s *Store) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // no file is not an error, just start with empty store
		}
		return fmt.Errorf("load: read file: %s %w", path, err)
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("load: unmarshal: %w", err)
	}

	for _, t := range tasks {
		s.tasks[t.ID] = t
		if t.ID >= s.nextID {
			s.nextID = t.ID + 1
		}
	}

	return nil

}
