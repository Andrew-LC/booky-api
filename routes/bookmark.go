package routes

import (
	"net/http"
	"bookmark-api/controller"
)

func BookmarkRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /create", controller.CreateBookmark)
	mux.HandleFunc("GET /all", controller.GetBookmarks)
	mux.HandleFunc("DELETE /delete/{id}", controller.DeleteBookmark)
	
	return mux
}
