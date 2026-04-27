package routes

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	TaskRoutes(r)
	UserRoutes(r)
}
