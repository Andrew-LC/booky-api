package controller

import (
	"errors"
	"net/http"
	"encoding/json"
	"bookmark-api/domain"
	"bookmark-api/service"
	middleware "bookmark-api/middlewares"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}

type AuthController struct {
	service service.AuthServiceInterface
}

func NewAuthController(s service.AuthServiceInterface) *AuthController {
	return &AuthController{service: s}
}

func (controller AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var creds domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := controller.service.CreateUser(creds)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidInput):
			WriteJSON(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, domain.ErrEmailExists):
			WriteJSON(w, http.StatusConflict, err.Error())

		default:
			WriteJSON(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	WriteJSON(w, http.StatusCreated, user)
}


func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var creds domain.SigninRequest

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		WriteJSON(w, http.StatusBadRequest, "invalid input")
		return
	}

	token, err := c.service.ValidateUser(creds)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidInput):
			WriteJSON(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, domain.ErrUserNotFound):
			WriteJSON(w, http.StatusNotFound, err.Error())

		case errors.Is(err, domain.ErrInvalidCredentials):
			WriteJSON(w, http.StatusUnauthorized, err.Error())

		default:
			WriteJSON(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (c *AuthController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		WriteJSON(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	err := c.service.DeleteUser(userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidInput):
			WriteJSON(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, domain.ErrUserNotFound):
			WriteJSON(w, http.StatusNotFound, err.Error())

		default:
			WriteJSON(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "account deleted",
	})
}
