package repo

import (
	"errors"
	"gorm.io/gorm"
	"bookmark-api/domain"
	"bookmark-api/model"
	"github.com/jackc/pgconn"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	DeleteUserAccount(userID uint) error
	
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (repo *UserRepo) CreateUser(user model.User) (model.User, error) {
	result := repo.db.Create(&user)

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

func (repo *UserRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *UserRepo) DeleteUserAccount(userID uint) error {
	result := repo.db.Delete(&model.User{}, userID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

