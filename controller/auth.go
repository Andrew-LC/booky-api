package controller

import (
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"bookmark-api/utils"
	"bookmark-api/model"
	"bookmark-api/domain"
	"bookmark-api/service"
	middleware "bookmark-api/middlewares"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
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


func Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user := model.GetUserByEmail(creds.Email)
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := model.DeleteUserAccount(userID); err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Account deleted")
}
