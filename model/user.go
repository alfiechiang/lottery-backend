package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string
	Password string
	Nickname string
	Status   int
	Avatar   string `gorm:"size:1000"`
}
