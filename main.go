package main

import (
	"taskmanager/routes"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()
	routes.SetUpRoutes(router)
	router.Run(":8000") // listens on 0.0.0.0:8000 by default
}
