package controllers

import (
	"net/http"
	"strconv"

	"brand-collab-tracker/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTaskHandler(c *gin.Context) {
	var input repositories.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	task, err := repositories.CreateTask(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTaskByProjectHandler(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	tasks, err := repositories.GetTaskByProject(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task list", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := repositories.GetTaskByID(uint(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input repositories.TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	task, err := repositories.UpdateTask(uint(id), input)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task successfull deleted"})
}