package main

import (
	"taskmanager/middlewares"
	"taskmanager/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	// global middlewates
	r.Use(middlewares.CORSMiddleware)

	routes.SetUpRoutes(r)
	r.Run(":8000")
}
