package model

import (
	"time"
	"bookmark-api/db"
)

type Account interface {
	CreateAccount() error
	DeleteAccount() error
} 

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

func DeleteUserAccount(userID uint) error {
	db := db.GetDB()
	var user User
	if err := db.Where("UserID = ?", userID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
