package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID uint
	gorm.Model
	UserName     string `gorm:"column:username"`
	Password     string
	Nickname     string
	Status       int
	Avatar       string `gorm:"size:1000"`
	Distributors string
	Role         string
}
