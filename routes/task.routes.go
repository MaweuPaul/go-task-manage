package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	r.POST("/tasks", controllers.CreateTaskHandler)
	r.GET("/tasks", controllers.GetAllTasksHandler)
	r.GET("/tasks/:id", controllers.GetTaskHandler)
	r.PUT("/tasks/:id", controllers.UpdateTaskHandler)
	r.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
}
