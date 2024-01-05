package models

import (
	"time"
)

type Users struct {
	Id        int        `json:"id" gorm:"unique;not null;column:id;"`
	Username  string     `json:"username" gorm:"unique;not null;column:username;" binding:"required,min=8,max=20" validate:"no_special_characters"`
	Password  string     `json:"password" gorm:"unique;not null;column:password;" binding:"required,min=8,max=40" validate:"no_special_characters"`
	Phone     string     `json:"phone" gorm:"column:phone;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
