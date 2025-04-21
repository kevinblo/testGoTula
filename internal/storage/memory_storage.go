package storage

import (
	"project/internal/models"
	"sync"
)

type MemoryStorage struct {
	mu   sync.RWMutex
	data map[string]models.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]models.Task)}
}

func (m *MemoryStorage) SaveTask(task models.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[task.ID] = task
	return nil
}

func (m *MemoryStorage) GetTask(id string) (models.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	task, exists := m.data[id]
	if !exists {
		return models.Task{}, nil
	}
	return task, nil
}
