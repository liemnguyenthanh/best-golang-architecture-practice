package models

import "time"

type Posts struct {
	Id        int        `json:"id"`
	User_id   int        `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type PostFilter struct {
	User_id int `json:"user_id"`
	Page    int `json:"page"`
	Limit   int `json:"limit"`
}
