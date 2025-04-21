package storage

import "project/internal/models"

type Storage interface {
	SaveTask(task models.Task) error
	GetTask(id string) (models.Task, error)
}
