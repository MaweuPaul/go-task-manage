package routes

import (
	"taskmanager/controllers"
	"taskmanager/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.Use(middlewares.RateLimitMiddleware("5-M"))
	{
		auth.POST("/register", controllers.CreateUserHandler)
		auth.POST("/login", controllers.LoginUserHandler)
	}

}
