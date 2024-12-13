package handlers

import (
	"github.com/gin-gonic/gin"
	"kf-golang-api/models"
	"kf-golang-api/services"
	"net/http"
)

// GetTasksHandler handles GET requests to fetch all tasks
func GetTasksHandler(c *gin.Context) {
	tasks := services.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

// AddTaskHandler handles POST requests to add a new task
func AddTaskHandler(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := services.AddTask(newTask)
	c.JSON(http.StatusCreated, task)
}
