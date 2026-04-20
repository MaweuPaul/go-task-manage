package main

import (
	"taskmanager/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	godotenv.Load()
	router := gin.Default()
	routes.SetUpRoutes(router)
	router.Run(":8000")
}
