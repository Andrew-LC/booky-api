package model

import (
	"time"
	"bookmark-api/db"
)

type Bookmark struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	URL       string    `gorm:"not null"`
	Title     string
	Notes     string
	Tags      []string  `gorm:"type:text[]"`
	CreatedAt time.Time
}

func (bookmark *Bookmark) CreateBookmark() error {
	database := db.GetDB()
	result := database.Create(bookmark)
	return result.Error
}

func GetBookmarks(userID uint) ([]Bookmark, error) {
	var bookmarks []Bookmark
	database := db.GetDB()
	result := database.Where("user_id = ?", userID).Find(&bookmarks)
	return bookmarks, result.Error
}

func DeleteBookmark(userID, id uint) error {
	database := db.GetDB()
	result := database.Where("id = ? AND user_id = ?", id, userID).Delete(&Bookmark{})
	return result.Error
}
