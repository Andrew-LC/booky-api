package routes

import (
	"net/http"
	"bookmark-api/controller"
)

func AuthRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /signup", controller.SignUp)
	mux.HandleFunc("POST /login", controller.Login)
	mux.HandleFunc("DELETE /deleteAcc", controller.DeleteAccount)
	mux.HandleFunc("GET /logout", controller.LogOut)
	
	return mux
}
