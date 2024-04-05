package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
)

func CreateCategory(c *gin.Context) {
	var body struct {
		UserID   uint   `json:"user_id" binding:"required"`
		Name     string `binding:"required"`
		IsIncome bool   `json:"is_income" binding:"omitempty"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid body type",
			"message": err.Error(),
		})
		return
	}

	var category = models.Category{
		Name:     body.Name,
		IsIncome: body.IsIncome,
		UserID:   body.UserID,
	}

	if err := initializers.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Category Created",
		"category": category,
	})
}

func GetCategories(c *gin.Context) {
	var categories []models.Category

	if err := initializers.DB.Preload("SubCategories").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get categories",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Categories Fetched!",
		"categories": categories,
	})
}

func GetCategory(c *gin.Context) {
	categoryId := c.Param("id")

	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing category id in path params",
		})
	}

	var category models.Category

	if err := initializers.DB.Preload("SubCategories").First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get category",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Category Fetched!",
		"category": category,
	})
}

func GetUserCategories(c *gin.Context) {
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing user_id in query params",
		})
		return
	}

	var categories []models.Category

	if err := initializers.DB.Preload("SubCategories").Where("user_id = ?", userId).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get categories",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Categories fetched!",
		"categories": categories,
	})
}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")

	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing category_id path param",
			"message": "Failed to delete category",
		})
		return
	}

	if err := initializers.DB.Delete(&models.Category{}, categoryId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":       "Failed to delete category",
			"message":     err.Error(),
			"category_id": categoryId,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted!",
	})
}

func UpdateCategory(c *gin.Context) {
	var body struct {
		Name     *string `binding:"omitempty"`
		IsIncome *bool   `json:"is_income" binding:"omitempty"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid body type",
			"message": err.Error(),
		})
		return
	}

	categoryId := c.Query("category_id")

	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing category_id query param",
			"message": "Failed to find category",
		})
		return
	}

	var category models.Category

	if err := initializers.DB.First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid Category ID",
			"message":    err.Error(),
			"account_id": categoryId,
		})
		return
	}

	if body.Name != nil {
		category.Name = *body.Name
	}
	if body.IsIncome != nil {
		category.IsIncome = *body.IsIncome
	}

	if err := initializers.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update category",
			"message": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Category updated successfully",
		"category": category,
	})
}
