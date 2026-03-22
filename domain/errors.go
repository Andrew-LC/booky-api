package domain

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrEmailExists  = errors.New("email already exists")
	ErrUserNotFound = errors.New("user not found")
)
