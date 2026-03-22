package repo

import (
	"errors"
	"context"
	"gorm.io/gorm"
	"bookmark-api/domain"
	"bookmark-api/model"
)

type BookmarkRepository interface {
	CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error)
	GetBookmarks(ctx context.Context, userID uint) ([]model.Bookmark, error)
	DeleteBookmark(ctx context.Context, userID, id uint) error
	UpdateBookmark(ctx context.Context, userID, id uint, updates map[string]interface{}) (*model.Bookmark, error)
}

type BookmarkRepo struct {
	db *gorm.DB
}

func NewBookmarkRepo(db *gorm.DB) BookmarkRepository {
	return &BookmarkRepo{db: db}
}

func (b *BookmarkRepo) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error) {
	if err := b.db.WithContext(ctx).Create(bookmark).Error; err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (b *BookmarkRepo) GetBookmarks(ctx context.Context, userID uint) ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark

	if err := b.db.WithContext(ctx).Where("user_id = ?", userID).Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (b *BookmarkRepo) DeleteBookmark(ctx context.Context, userID, id uint) error {
	result := b.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.Bookmark{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrBookmarkNotFound
	}

	return nil
}

func (b *BookmarkRepo) UpdateBookmark(ctx context.Context, userID, id uint, updates map[string]interface{}) (*model.Bookmark, error) {
	var bookmark model.Bookmark

	if err := b.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).
		First(&bookmark).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrBookmarkNotFound
		}
		return nil, err
	}

	if len(updates) == 0 {
		return &bookmark, nil
	}

	if err := b.db.Model(&bookmark).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}
