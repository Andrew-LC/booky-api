package model

import (
    "time"
    "bookmark-api/db"
    "gorm.io/gorm"
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

func UpdateBookmark(userID, id uint, updates map[string]interface{}) (*Bookmark, error) {
    database := db.GetDB()

    if len(updates) == 0 {
        var current Bookmark
        if err := database.Where("id = ? AND user_id = ?", id, userID).First(&current).Error; err != nil {
            return nil, err
        }
        return &current, nil
    }

    result := database.Model(&Bookmark{}).
        Where("id = ? AND user_id = ?", id, userID).
        Updates(updates)
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected == 0 {
        return nil, gorm.ErrRecordNotFound
    }

    var updated Bookmark
    if err := database.Where("id = ? AND user_id = ?", id, userID).First(&updated).Error; err != nil {
        return nil, err
    }
    return &updated, nil
}
