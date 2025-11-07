package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"brand-collab-tracker/models"
	"brand-collab-tracker/repositories"
)

func CreateCategoryHandler(c *gin.Context) {
	var category models.CategoryMaster
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := repositories.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make category", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func GetCategoriesHandler(c *gin.Context) {
	categories, err := repositories.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category list", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func UpdateCategoryHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := repositories.GetCategoryByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := repositories.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category successfully deleted"})
}