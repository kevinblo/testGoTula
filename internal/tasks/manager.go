package tasks

import (
	"fmt"
	"time"

	"project/internal/models"
	"project/internal/storage"
)

func StartTaskProcessing(task models.Task, store storage.Storage) {
	time.Sleep(3 * time.Minute)
	task.Status = "completed"
	task.Result = fmt.Sprintf("Результат задачи %s", task.ID)
	err := store.SaveTask(task)
	if err != nil {
		return
	}
}
