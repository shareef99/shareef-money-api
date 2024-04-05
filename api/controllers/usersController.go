package controllers

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
	"gorm.io/gorm"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func Signin(c *gin.Context) {
	var body struct {
		Name  string `binding:"required"`
		Email string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User doesn't exist",
				"error":   err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to signin",
				"error":   err.Error(),
			})
			return
		}
	}

	if user.ID == 0 {
		referCode := generateRandomString(6)

		newUser := models.User{
			Name:      body.Name,
			Email:     body.Email,
			ReferCode: referCode,
		}

		if err := initializers.DB.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create user",
				"error":   err.Error(),
			})
			return
		}

		var createdUser models.User

		if err := initializers.DB.First(&createdUser, newUser.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to fetch created user",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User Signup successfully",
			"user":    createdUser,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Signin successfully",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {

	var users []models.User

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

	var user models.User
	err := initializers.DB.Preload("Accounts").Preload("Categories").Preload("Transactions").First(&user, id)

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
		Mobile         *string `binding:"omitempty"`
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

	user := models.User{
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

	deleteUser := initializers.DB.Delete(&models.User{}, id)

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

func UpdateUser(c *gin.Context) {
	var body struct {
		Name           *string `binding:"omitempty"`
		Mobile         *string `binding:"omitempty"`
		Currency       *string `binding:"omitempty"`
		MonthStartDate *uint8  `json:"month_start_date" binding:"omitempty"`
		WeekStartDay   *string `json:"week_start_day" binding:"omitempty,oneof=mon tue wed thu fri sat sun"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid body type",
			"message": err.Error(),
		})
		return
	}

	id := c.Query("id")

	var user models.User
	err := initializers.DB.First(&user, id)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid User ID",
			"message": err.Error.Error(),
			"id":      id,
		})
		return
	}

	if body.Name != nil {
		user.Name = *body.Name
	}

	if body.Mobile != nil {
		if *body.Mobile == "null" {
			user.Mobile = nil
		} else {
			user.Mobile = body.Mobile
		}
	}

	if body.Currency != nil {
		user.Currency = *body.Currency
	}
	if body.MonthStartDate != nil {
		user.MonthStartDate = *body.MonthStartDate
	}
	if body.WeekStartDay != nil {
		user.WeekStartDay = *body.WeekStartDay
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update user",
			"message": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}
