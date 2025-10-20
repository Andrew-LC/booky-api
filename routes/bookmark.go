package routes

import (
	"net/http"
	"bookmark-api/controller"
)

func BookmarkRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /", controller.CreateBookmark)
	mux.HandleFunc("GET /", controller.GetBookmarks)
	mux.HandleFunc("DELETE /{id}", controller.DeleteBookmark)
	mux.HandleFunc("PUT /{id}", controller.UpdateBookmark)
	
	return mux
}
