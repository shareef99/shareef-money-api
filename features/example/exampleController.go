package example

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/initializers"
)

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Example routes under v1!",
	})
}

func CreateExample(c *gin.Context) {
	// GET DATA OFF REQ BODY
	var body struct {
		Name  string `binding:"required"`
		Email string `binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println("Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind data",
			"error":   err.Error(),
		})
		return
	}

	// CREATE RECORD

	example := Example{
		Name:  "test",
		Email: "test@test.com",
	}

	result := initializers.DB.Create(&example)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create example",
			"error":   err.Error(),
		})
		return
	}

	// RETURN IT
	c.JSON(http.StatusCreated, gin.H{
		"message": "Example created!",
		"example": example,
	})
}
