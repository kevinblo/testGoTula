package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskResult).Methods("GET")

	return router
}
