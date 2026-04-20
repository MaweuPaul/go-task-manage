package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.CreateUserHandler)
		auth.POST("/login", controllers.LoginUserHandler)
	}

}
