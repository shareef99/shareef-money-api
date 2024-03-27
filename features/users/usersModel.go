package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string
	Email          string
	Mobile         *string
	Currency       string
	MonthStartDate string
	WeekStartDay   string `gorm:"type:varchar(3);not null;check(week_start_day in ('mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun'))"`
	ReferCode      string
}
