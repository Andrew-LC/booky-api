package routes

import (
	"net/http"

	"gorm.io/gorm"
	"bookmark-api/repo"
	"bookmark-api/service"
	"bookmark-api/controller"
	"bookmark-api/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(mux *http.ServeMux, db *gorm.DB) {
	authRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)
	authMux := AuthRoutes(authController)
	bookmarkMux := BookmarkRoutes()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/bookmark/",
		http.StripPrefix("/bookmark",
			middleware.JWTMiddleware(bookmarkMux),
		),
	)
	mux.Handle("/metrics", promhttp.Handler())
}
