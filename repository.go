package main

import (
	"encoding/json"
	"log/slog"
)

func CreateTask(task Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		slog.Error("Could not marshal task.", "error", err)
		return err
	}

	// SET → responsible for storing the task DATA in Redis
	// Here we save the full task object (as JSON) under a unique key
	// Example:
	// key:   task:123
	// value: {"id":"123","title":"Study Go","done":false}
	//
	// The "0" means no expiration (no TTL), so the key persists indefinitely
	err = rdb.Set(ctx, "task-"+task.UUID, data, 0).Err()
	if err != nil {
		slog.Error("Could not save task.", "error", err)
		return err
	}

	// RPUSH → responsible for organizing/registering task IDs in a list
	// Here we DO NOT store the task data, only its ID
	//
	// This list acts as a "manual index", since Redis does not provide
	// an easy way to list all keys efficiently
	// Example:
	// key: tasks
	// value: ["123", "456", "789"]
	//
	// This allows us to later iterate over all IDs and fetch each task
	// individually using GET task:{id}
	rdb.RPush(ctx, "tasks", task.UUID)

	return nil
}

func GetAllTasks() ([]Task, error) {
	indexesId, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		slog.Error("Could not get indexes.", "error", err)
		return []Task{}, err
	}

	tasks := []Task{}
	for _, id := range indexesId {

		data, err := rdb.Get(ctx, "task-"+id).Result()
		if err != nil {
			slog.Error("Could not get a task.", "error", err)
			return []Task{}, err
		}

		task := Task{}
		json.Unmarshal([]byte(data), &task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
