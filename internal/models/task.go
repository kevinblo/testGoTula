package models

type Task struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Result    string `json:"result,omitempty"`
	CreatedAt string `json:"created_at"`
}
