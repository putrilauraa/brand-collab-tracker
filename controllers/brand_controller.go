package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"brand-collab-tracker/repositories"
	"gorm.io/gorm"
)

func CreateBrandHandler(c *gin.Context) {
	var input repositories.BrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	brand, err := repositories.CreateBrand(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand)
}

func GetBrandsHandler(c *gin.Context) {
	brands, err := repositories.GetBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brand list", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brands)
}

func GetBrandByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	brand, err := repositories.GetBrandByID(uint(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func UpdateBrandHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input repositories.BrandInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	brand, err := repositories.UpdateBrand(uint(id), input)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func DeleteBrandHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteBrand(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand sucessfully deleted"})
}