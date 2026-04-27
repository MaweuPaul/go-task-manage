package routes

import (
	"taskmanager/controllers"
	"taskmanager/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	protected := r.Group("/tasks")
	protected.Use(middlewares.RateLimitMiddleware("200-H"))
	protected.Use(
		middlewares.AuthMiddleware,
	)

	protected.POST("", controllers.CreateTaskHandler)
	protected.GET("", controllers.GetAllTasksHandler)
	protected.GET("/:id", controllers.GetTaskHandler)
	protected.PUT("/:id", controllers.UpdateTaskHandler)
	protected.DELETE("/:id", controllers.DeleteTaskHandler)
}
