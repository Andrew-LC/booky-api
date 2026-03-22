package model

import "time"

type Bookmark struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;constraint:OnDelete:CASCADE"`
	URL       string    `gorm:"not null"`
	Title     string
	Notes     string
	Image     string    
	Tags      []string  `gorm:"type:text[]"`
	CreatedAt time.Time
}

