package services

import (
	"time"

	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
)

type DailyTransaction struct {
	TransactionAt time.Time            `json:"transaction_at"`
	TotalIncome   float32              `json:"total_income"`
	TotalExpense  float32              `json:"total_expense"`
	Transactions  []models.Transaction `json:"transactions"`
}

func GetDailyTransactionByMonth(month string, userID string) ([]DailyTransaction, error) {
	var transactions []models.Transaction

	if err := initializers.DB.
		Preload("Category").
		Preload("Account").
		Where("user_id = ? AND EXTRACT(MONTH FROM transaction_at) = ?", userID, month).
		Order("transaction_at ASC").
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	dailyTransactionsMap := make(map[time.Time]DailyTransaction)
	for _, transaction := range transactions {
		transactionAt := transaction.TransactionAt.Truncate(24 * time.Hour)
		dailyTransaction, ok := dailyTransactionsMap[transactionAt]
		if !ok {
			dailyTransaction = DailyTransaction{
				TransactionAt: transactionAt,
			}
		}

		if transaction.Type == "income" {
			dailyTransaction.TotalIncome += transaction.Amount
		} else if transaction.Type == "expense" {
			dailyTransaction.TotalExpense += transaction.Amount
		}

		dailyTransaction.Transactions = append(dailyTransaction.Transactions, transaction)
		dailyTransactionsMap[transactionAt] = dailyTransaction
	}

	var dailyTransactions []DailyTransaction
	for _, dailyTransaction := range dailyTransactionsMap {
		dailyTransactions = append(dailyTransactions, dailyTransaction)
	}

	return dailyTransactions, nil
}
