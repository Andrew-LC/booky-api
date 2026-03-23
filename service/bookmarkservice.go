package service

import (
	"context"
	"bookmark-api/model"
	"bookmark-api/repo"
)

type BookmarkServiceInterface interface {
	CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error)
	GetBookmarks(ctx context.Context, userID uint) ([]model.Bookmark, error)
	DeleteBookmark(ctx context.Context, userID, id uint) error
	UpdateBookmark(ctx context.Context, userID, id uint, updates map[string]interface{}) (*model.Bookmark, error)
	UpdateBookmarkMeta(ctx context.Context, id uint, title, image string) error
}

type BookmarkService struct {
	repo repo.BookmarkRepository
}

func NewBookmarkService(r repo.BookmarkRepository) BookmarkServiceInterface {
	return &BookmarkService{repo: r}
}

func (b *BookmarkService) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) (*model.Bookmark, error) {
	return b.repo.CreateBookmark(ctx, bookmark)
}

func (b *BookmarkService) GetBookmarks(ctx context.Context, userID uint) ([]model.Bookmark, error) {
	return b.repo.GetBookmarks(ctx, userID)
}

func (b *BookmarkService) DeleteBookmark(ctx context.Context, userID, id uint) error {
	return b.repo.DeleteBookmark(ctx, userID, id)
}

func (b *BookmarkService) UpdateBookmark(ctx context.Context, userID, id uint, updates map[string]interface{}) (*model.Bookmark, error) {
	return b.repo.UpdateBookmark(ctx, userID, id, updates)
}

func (b *bookmarkService) UpdateBookmarkMeta(ctx context.Context, id uint, title, image string) error {
	return b.repo.UpdateMeta(ctx, id, title, image)
}
