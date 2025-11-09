package controllers

import (
	"net/http"
	"strconv"

	"brand-collab-tracker/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProjectHandler(c *gin.Context) {
	var input repositories.ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	project, err := repositories.CreateProject(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func GetProjectsHandler(c *gin.Context) {
	projects, err := repositories.GetBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project list", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func GetProjectByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	project, err := repositories.GetProjectByID(uint(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func UpdateProjectHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input repositories.ProjectInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	project, err := repositories.UpdateProject(uint(id), input)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func DeleteProjectHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteProject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project successfull deleted"})
}