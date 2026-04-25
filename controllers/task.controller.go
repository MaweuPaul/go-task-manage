package controllers

import (
	"taskmanager/models"
	"taskmanager/services"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := services.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Task created successfully", "task": createdTask})
}

func UpdateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updateTask, err := services.UpdateTask(c.Param("id"), task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task updated successfully", "task": updateTask})
}

func DeleteTaskHandler(c *gin.Context) {
	if err := services.DeleteTask(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func GetTaskHandler(c *gin.Context) {

	if task, err := services.GetTask(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"task": task})
	}
}

func GetAllTasksHandler(c *gin.Context) {

	if tasks, err := services.GetAllTasks(); err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"task": tasks})
	}
}
