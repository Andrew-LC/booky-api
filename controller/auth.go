package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bookmark-api/utils"
	"bookmark-api/model"
        "golang.org/x/crypto/bcrypt"
	middleware "bookmark-api/middlewares"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password),  bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server Error",  http.StatusInternalServerError)
	}

	user := model.User{
		Username: utils.GenerateUsername(creds.Email),
		Email: creds.Email,
		Password: string(hashedPassword),
	}

	if err := user.CreateUser(); err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
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

func LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "loggin you out")
}
