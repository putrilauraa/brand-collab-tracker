package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"brand-collab-tracker/repositories"
	"gorm.io/gorm"
)

func CreateAttachmentHandler(c *gin.Context) {
	var input repositories.AttachmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	attachment, err := repositories.CreateAttachment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attachment)
}

func GetAttachmentsByProjectHandler(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	attachments, err := repositories.GetAttachmentsByProject(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attachment list", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

func GetAttachmentByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	attachment, err := repositories.GetAttachmentByID(uint(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attachment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

func UpdateAttachmentHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input repositories.AttachmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attachment input", "details": err.Error()})
		return
	}

	attachment, err := repositories.UpdateAttachment(uint(id), input)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attachment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

func DeleteAttachmentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteAttachment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attachment successfully deleted"})
}