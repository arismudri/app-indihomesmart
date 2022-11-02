package models

import (
	"time"
)

type Article struct {
	Id        int        `json:"id"`
	Title     string     `json:"title" binding:"required,min=20,max=200" gorm:"type:varchar(200)"`
	Content   string     `json:"content" binding:"required,min=200,max=250" gorm:"type:text"`
	Category  string     `json:"category" binding:"required,min=3,max=100" gorm:"type:varchar(100)"`
	Status    string     `json:"status" binding:"required,oneof=publish draft thrash" gorm:"type:varchar(100)"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

// TableName is Database TableName of this model
func (e *Article) TableName() string {
	return "articles"
}
