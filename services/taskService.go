package services

import "kf-golang-api/models"

var tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Details: "Start by building a simple API"},
}

// GetTasks
/**
Returns all tasks
*/
func GetTasks() []models.Task {
	return tasks
}

// AddTask
/**
Adds a new task
*/
func AddTask(task models.Task) models.Task {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
	return task
}
