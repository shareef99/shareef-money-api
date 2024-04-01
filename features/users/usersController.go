package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/initializers"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {

	var users []User

	if err := initializers.DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Users not found",
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to get users",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched!",
		"users":   users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Query("id")

	var user User
	err := initializers.DB.First(&user, id)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to get user",
			"message": err.Error.Error(),
			"id":      id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User found!",
		"user":    user,
	})
}

func CreateUser(c *gin.Context) {
	var body struct {
		Name           string  `binding:"required"`
		Email          string  `binding:"required,email"`
		Mobile         string  `binding:"required"`
		Currency       *string `binding:"omitempty"`
		MonthStartDate *uint8  `json:"month_start_date" binding:"omitempty"`
		WeekStartDay   *string `json:"week_start_day" binding:"omitempty,oneof=mon tue wed thu fri sat sun"`
		ReferCode      string  `json:"refer_code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to bind data",
			"message": err.Error(),
		})
		return
	}

	user := User{
		Name:      body.Name,
		Email:     body.Email,
		Mobile:    body.Mobile,
		ReferCode: body.ReferCode,
	}

	if body.WeekStartDay != nil {
		user.WeekStartDay = *body.WeekStartDay
	}

	if body.Currency != nil {
		user.Currency = *body.Currency
	}

	if body.MonthStartDate != nil {
		user.MonthStartDate = *body.MonthStartDate
	}

	result := initializers.DB.Create(&user)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deleteUser := initializers.DB.Delete(&User{}, id)

	if deleteUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to delete user",
			"message": deleteUser.Error.Error(),
			"id":      id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted!",
	})
}
