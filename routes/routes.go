package routes

import (
	"net/http"

	"bookmark-api/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(mux *http.ServeMux) {
	authMux := AuthRoutes()
	bookmarkMux := BookmarkRoutes()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/bookmark/", http.StripPrefix("/bookmark", middleware.JWTMiddleware(bookmarkMux)))
	mux.Handle("/metrics", promhttp.Handler())
}
