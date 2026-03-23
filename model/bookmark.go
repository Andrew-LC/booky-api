package model

import (
	"time"
	"github.com/lib/pq"
)

type Bookmark struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;constraint:OnDelete:CASCADE"`
	URL       string    `gorm:"not null"`
	Title     string
	Notes     string
	Image     string    
	Tags   pq.StringArray `gorm:"type:text[]"`  //
	CreatedAt time.Time
}

