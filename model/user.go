package model

import (
	"time"
	"bookmark-api/db"
)

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"unique;not null"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	CreatedAt time.Time
	Bookmarks []Bookmark `gorm:"foreignKey:UserID"`
}

func (user *User) CreateUser() error {
	db := db.GetDB() 
	result := db.Create(user)

	return result.Error
}

func GetUserByEmail(email string) *User {
	db := db.GetDB()
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
