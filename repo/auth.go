package repo

import (
	"errors"
	"context"
	"gorm.io/gorm"
	"bookmark-api/domain"
	"bookmark-api/model"
	"github.com/jackc/pgconn"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	DeleteUserAccount(ctx context.Context, userID uint) error
	
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	result := repo.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			switch pgErr.Code {
			case "23505": 
				return model.User{}, domain.ErrEmailExists
			}
		}
		return model.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, domain.ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepo) DeleteUserAccount(ctx context.Context, userID uint) error {
	result := repo.db.WithContext(ctx).Delete(&model.User{}, userID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

