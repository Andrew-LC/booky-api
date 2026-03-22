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
	ValidateUser(signinrequest domain.SigninRequest) (string, error)
	DeleteUser(userID uint) error
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


func (authservice *AuthService) ValidateUser(signinrequest domain.SigninRequest) (string, error) {
	if signinrequest.Email == "" || signinrequest.Password == "" {
		return "", domain.ErrInvalidInput
	}

	user, err := authservice.repo.GetUserByEmail(signinrequest.Email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(signinrequest.Password),
	); err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authservice *AuthService) DeleteUser(userID uint) error {
	if userID == 0 {
		return domain.ErrInvalidInput
	}

	err := authservice.repo.DeleteUserAccount(userID)
	if err != nil {
		return err
	}

	return nil
}
