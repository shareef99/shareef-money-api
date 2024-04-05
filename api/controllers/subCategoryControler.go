package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
)

func CreateSubCategory(c *gin.Context) {
	var body struct {
		Name       string `json:"name" binding:"required"`
		CategoryID uint   `json:"category_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   "Invalid payload",
			"message": err.Error(),
		})
	}

	subCategory := models.SubCategory{
		Name:       body.Name,
		CategoryID: body.CategoryID,
	}

	if err := initializers.DB.Create(&subCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create sub category",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Sub Category Created",
		"sub_category": subCategory,
	})
}

func GetSubCategory(c *gin.Context) {
	subCategoryId := c.Param("id")

	if subCategoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing sub category id in path params",
		})
	}

	var subCategory models.SubCategory

	if err := initializers.DB.First(&subCategory, subCategoryId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get sub category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Sub Category Fetched!",
		"sub_category": subCategory,
	})
}

func GetSubCategories(c *gin.Context) {
	var subCategories []models.SubCategory

	if err := initializers.DB.Find(&subCategories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch SubCategories",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Fetched Sub Categories",
		"sub_categories": subCategories,
	})
}

func UpdateSubCategory(c *gin.Context) {
	subCategoryId := c.Param("id")

	if subCategoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing sub category id in path param",
		})
		return
	}

	var body struct {
		Name string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid payload",
			"message": err.Error(),
		})
		return
	}

	var subCategory models.SubCategory

	if err := initializers.DB.First(&subCategory, subCategoryId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Invalid Sub Category ID",
			"message":         err.Error(),
			"sub_category_id": subCategoryId,
		})
		return
	}

	subCategory.Name = body.Name

	if err := initializers.DB.Save(&subCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update sub category",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Sub Category Updated",
		"sub_category": subCategory,
	})
}

func DeleteSubCategory(c *gin.Context) {
	subCategoryId := c.Param("id")

	if subCategoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing sub category id in path param",
		})
		return
	}

	if err := initializers.DB.Delete(&models.SubCategory{}, subCategoryId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":           "Failed to delete sub category",
			"message":         err.Error(),
			"sub_category_id": subCategoryId,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sub Category deleted!",
	})
}
