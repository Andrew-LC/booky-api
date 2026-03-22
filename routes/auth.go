package routes

import (
	"net/http"
	"bookmark-api/controller"
)

func AuthRoutes(authController *controller.AuthController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /signup", authController.SignUp)
	mux.HandleFunc("POST /login", authController.Login)
	mux.HandleFunc("DELETE /deleteAcc", authController.DeleteAccount)

	return mux
}
