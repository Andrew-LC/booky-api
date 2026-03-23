package routes

import (
	"net/http"

	"bookmark-api/metrics"
	"gorm.io/gorm"
	"bookmark-api/repo"
	"bookmark-api/service"
	"bookmark-api/controller"
	"bookmark-api/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(mux *http.ServeMux, db *gorm.DB) http.Handler {
	metrics.Init()

	authRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)
	authMux := AuthRoutes(authController)

	bookmarkRepo := repo.NewBookmarkRepo(db)
	bookmarkService := service.NewBookmarkService(bookmarkRepo)
	bookmarkController := controller.NewBookmarkController(bookmarkService)
	bookmarkMux := BookmarkRoutes(bookmarkController)

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/bookmark/",
		http.StripPrefix("/bookmark",
			middleware.JWTMiddleware(bookmarkMux),
		),
	)
	mux.Handle("/metrics", promhttp.Handler())

	return middleware.MetricsMiddleware(middleware.LoggingMiddleware(mux))
}
