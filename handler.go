package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createTaskHandler(w, r)
	case http.MethodDelete:
		getTasksHandler(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskRequest := CreateTaskRequest{}

	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		slog.Error("Coud not read task request body.", "error", err)
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	newTask := toTask(taskRequest)
	err = CreateTask(newTask)
	if err != nil {
		slog.Error("Could not create task.", "error", err)
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}
}

func toTask(tr CreateTaskRequest) Task {
	return Task{
		UUID:  uuid.New().String(),
		Title: tr.Title,
		Done:  tr.Done,
	}
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := GetAllTasks()
	if err != nil {
		http.Error(w, "failed to get taks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
