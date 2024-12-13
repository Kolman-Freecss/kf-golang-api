package routes

import (
	"github.com/gin-gonic/gin"
	"kf-golang-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasksHandler)
	r.POST("/tasks", handlers.AddTaskHandler)

	return r
}
