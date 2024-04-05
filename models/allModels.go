package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string     `json:"name" gorm:"not null"`
	Email          string     `json:"email" gorm:"unique;not null"`
	Mobile         *string    `json:"mobile" gorm:"unique"`
	Currency       string     `json:"currency" gorm:"default:inr"`
	MonthStartDate uint8      `json:"month_start_date" gorm:"not null;default:1"`
	WeekStartDay   string     `json:"week_start_day" gorm:"not null;default:mon"`
	ReferCode      string     `json:"refer_code" gorm:"unique;not null"`
	Accounts       []Account  `json:"accounts" gorm:"foreignKey:UserID;constraints:OnDelete:SET NULL"`
	Categories     []Category `json:"categories" gorm:"foreignKey:UserID;constraints:OnDelete:SET NULL"`
}

type Account struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Amount      float32 `json:"amount" gorm:"not null"`
	Description *string `json:"description"`
	IsHidden    bool    `json:"is_hidden" gorm:"default:false"`
	UserID      uint    `json:"user_id" gorm:"not null"`
}

type Category struct {
	gorm.Model
	Name          string        `json:"name" gorm:"not null"`
	IsIncome      bool          `json:"is_income" gorm:"not null"`
	UserID        uint          `json:"user_id" gorm:"not null"`
	SubCategories []SubCategory `json:"sub_categories" gorm:"foreignKey:CategoryID;constraints:OnDelete:SET NULL"`
}

type SubCategory struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	CategoryID uint   `json:"category_id" gorm:"not null"`
}
