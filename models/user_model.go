package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" binding:"required,min=3,max=200" gorm:"type:varchar(200)"`
	Email     string     `json:"email" binding:"required,min=6,max=200,email" gorm:"type:varchar(200);index:idx_email,unique"`
	Username  string     `json:"username" binding:"required,max=200,alphanum" gorm:"type:varchar(200)"`
	Password  string     `json:"password,omitempty" binding:"required,min=8,max=200" gorm:"type:varchar(200)"`
	Bio       string     `json:"bio" binding:"required,max=200" gorm:"type:varchar(200)"`
	IsLogin   uint16     `json:"is_login,omitempty" gorm:"type:tinyint"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

// TableName is Database TableName of this model
func (e *User) TableName() string {
	return "users"
}
