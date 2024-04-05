package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
	"gorm.io/gorm"
)

func GetAccounts(c *gin.Context) {
	var accounts []models.Account

	if err := initializers.DB.Find(&accounts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Accounts not found",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to get accounts",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Accounts fetched!",
		"accounts": accounts,
	})
}

func GetUserAccounts(c *gin.Context) {
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Send user_id",
		})
	}

	var user models.User

	if err := initializers.DB.Preload("Accounts").First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Accounts not found",
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to get user accounts",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User accounts fetched!",
		"accounts": user.Accounts,
	})
}

func CreateAccount(c *gin.Context) {
	var body struct {
		Name        string  `binding:"required"`
		Amount      float32 `binding:"required"`
		Description *string `binding:"omitempty"`
		IsHidden    bool    `json:"is_hidden" binding:"required"`
		UserId      uint    `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to bind data",
			"message": err.Error(),
		})
		return
	}

	account := models.Account{
		Name:        body.Name,
		Amount:      body.Amount,
		Description: body.Description,
		IsHidden:    body.IsHidden,
		UserID:      body.UserId,
	}

	if body.Description != nil {
		if *body.Description == "null" {
			account.Description = nil
		}
	}

	result := initializers.DB.Create(&account)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create account",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created!",
		"account": account,
	})
}

func DeleteAccount(c *gin.Context) {
	accountId := c.Param("id")

	if accountId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing account_id path param",
			"message": "Failed to delete account",
		})
		return
	}

	if err := initializers.DB.Delete(&models.Account{}, accountId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Failed to delete account",
			"message":    err.Error(),
			"account_id": accountId,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted!",
	})
}

func UpdateAccount(c *gin.Context) {
	var body struct {
		Name        *string  `binding:"omitempty"`
		Amount      *float32 `binding:"omitempty"`
		Description *string  `binding:"omitempty"`
		IsHidden    *bool    `json:"is_hidden" binding:"omitempty"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid body type",
			"message": err.Error(),
		})
		return
	}

	accountId := c.Query("account_id")

	if accountId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing account_id query param",
			"message": "Failed to find account",
		})
		return
	}

	var account models.Account

	if err := initializers.DB.First(&account, accountId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid Account ID",
			"message":    err.Error(),
			"account_id": accountId,
		})
		return
	}

	if body.Name != nil {
		account.Name = *body.Name
	}

	if body.Description != nil {
		if *body.Description == "null" {
			account.Description = nil
		} else {
			account.Description = body.Description
		}
	}

	if body.Amount != nil {
		account.Amount = *body.Amount
	}
	if body.IsHidden != nil {
		account.IsHidden = *body.IsHidden
	}

	if err := initializers.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update account",
			"message": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account updated successfully",
		"account": account,
	})
}
