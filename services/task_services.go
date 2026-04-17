package services

import (
	"fmt"
	"taskmanager/models"

	"github.com/google/uuid"
)

// in memory task storage
var tasks = make(map[string]models.Task)

func CreateTask(task models.Task) (models.Task, error) {

	task.ID = uuid.New().String()

	tasks[task.ID] = task

	return task, nil
}

// update the task with the given id

func UpdateTask(id string, task models.Task) (models.Task, error) {

	if _, exists := tasks[id]; !exists {
		return models.Task{}, fmt.Errorf("task not found")
	}

	// keep the same id
	task.ID = id
	// update the task
	tasks[id] = task

	return task, nil

}

func DeleteTask(id string) error {

	if _, exists := tasks[id]; !exists {
		return fmt.Errorf("Task not found")

	}
	delete(tasks, id)
	return nil
}

func GetTask(id string) (models.Task, error) {

	retrievedTask, exists := tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("Task not found")
	}

	return retrievedTask, nil
}

func GetAllTasks() ([]models.Task, error) {

	var tasksList []models.Task

	for _, t := range tasks {
		tasksList = append(tasksList, t)
	}

	return tasksList, nil
}

// delete all tasks

func DeleteAllTasks() error {

	tasks = make(map[string]models.Task)
	return nil
}

// delete multiple tasks by ids
func DeleteMultipleTasks(ids []string) error {

	var notFoundIds []string

	for _, id := range ids {
		if _, exists := tasks[id]; !exists {
			notFoundIds = append(notFoundIds, id)
		}
	}

	if len(notFoundIds) > 0 {
		return fmt.Errorf("tasks not found: %v", notFoundIds)
	}

	for _, id := range ids {
		delete(tasks, id)
	}
	return nil
}
