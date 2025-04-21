package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"project/internal/models"
	"project/internal/storage"
	"project/internal/tasks"
)

var taskStore storage.Storage

func SetupStorage(s storage.Storage) {
	taskStore = s
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	task := models.Task{
		ID:        id,
		Status:    "processing",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	taskStore.SaveTask(task)
	go tasks.StartTaskProcessing(task, taskStore)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"task_id": id})
}

func GetTaskResult(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, err := taskStore.GetTask(id)
	if err != nil || task.ID == "" {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}
