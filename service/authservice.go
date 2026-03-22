package service

import (
	"bookmark-api/utils"
	"bookmark-api/repo"
	"bookmark-api/model"
	"bookmark-api/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	CreateUser(signuprequest domain.SignupRequest) (model.User, error)
}

type AuthService struct {
	repo repo.UserRepository
}

func NewAuthService(repo repo.UserRepository) AuthServiceInterface {
	return &AuthService{repo: repo}
}

func (authservice *AuthService) CreateUser(signuprequest domain.SignupRequest) (model.User, error) {
	if signuprequest.Email == "" || signuprequest.Password == "" {
		return model.User{}, domain.ErrInvalidInput
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(signuprequest.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: utils.GenerateUsername(signuprequest.Email),
		Email:    signuprequest.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := authservice.repo.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	createdUser.Password = "" 
	return createdUser, nil
}
