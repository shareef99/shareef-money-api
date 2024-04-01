package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	Mobile         string `json:"mobile" gorm:"unique"`
	Currency       string `json:"currency" gorm:"default:inr"`
	MonthStartDate uint8  `json:"month_start_date" gorm:"not null;default:1"`
	WeekStartDay   string `json:"week_start_day" gorm:"not null;default:mon"`
	ReferCode      string `json:"refer_code" gorm:"unique;not null"`
}
