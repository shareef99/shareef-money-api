package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/api/services"
	"github.com/shareef99/shareef-money-api/initializers"
)

func CreateTransaction(c *gin.Context) {
	var body struct {
		UserID        uint                   `json:"user_id" binding:"required"`
		AccountID     uint                   `json:"account_id" binding:"required"`
		CategoryID    uint                   `json:"category_id" binding:"required"`
		Type          models.TransactionType `json:"type" binding:"required"`
		Notes         *string                `json:"notes" binding:"omitempty"`
		Amount        float32                `json:"amount" binding:"required"`
		TransactionAt time.Time              `json:"transaction_at" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to bind data",
			"message": err.Error(),
		})
		return
	}

	transaction := models.Transaction{
		UserID:        body.UserID,
		AccountID:     body.AccountID,
		CategoryID:    body.CategoryID,
		Type:          body.Type,
		Notes:         body.Notes,
		Amount:        body.Amount,
		TransactionAt: body.TransactionAt,
	}

	if body.Notes != nil {
		if *body.Notes == "null" {
			transaction.Notes = nil
		}
	}

	if err := initializers.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create transaction",
			"message": err.Error(),
		})
		return
	}

	account := models.Account{}

	if err := initializers.DB.First(&account, body.AccountID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get account",
			"message": err.Error(),
		})
		return
	}

	account.Amount -= body.Amount

	if err := initializers.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update account balance",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Transaction created!",
		"transaction": transaction,
	})
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	if err := initializers.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get transactions",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions Fetched!",
		"transactions": transactions,
	})
}

func GetAccountTransactions(c *gin.Context) {
	accountId := c.Param("id")

	if accountId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing account id in path param",
		})
		return
	}

	var transactions []models.Transaction

	if err := initializers.DB.Where("account_id = ?", accountId).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get transactions",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions Fetched!",
		"transactions": transactions,
	})
}

func GetUserTransactions(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing user id in path param",
		})
		return
	}

	var transactions []models.Transaction

	if err := initializers.DB.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get transactions",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions Fetched!",
		"transactions": transactions,
	})
}

func GetCategoryTransactions(c *gin.Context) {
	categoryId := c.Param("id")

	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing category id in path param",
		})
		return
	}

	var transactions []models.Transaction

	if err := initializers.DB.Where("category_id = ?", categoryId).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get transactions",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions Fetched!",
		"transactions": transactions,
	})
}

func GetMonthlyTransactions(c *gin.Context) {
	userID := c.Query("user_id")
	month := c.Query("month")

	if userID == "" || month == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing user_id or month in query params",
		})
		return
	}

	dailyTransactions, err := services.GetDailyTransactionByMonth(month, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get transactions",
			"message": err.Error(),
		})
		return
	}

	if len(dailyTransactions) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":      "No transactions found for this month",
			"transactions": []services.DailyTransaction{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions Fetched!",
		"transactions": dailyTransactions,
	})
}
