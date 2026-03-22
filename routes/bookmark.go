package routes

import (
	"net/http"
	"bookmark-api/controller"
)

func BookmarkRoutes(bookmarkController *controller.BookmarkController) *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /", bookmarkController.CreateBookmark)
	mux.HandleFunc("GET /", bookmarkController.GetBookmarks)
	mux.HandleFunc("DELETE /{id}", bookmarkController.DeleteBookmark)
	mux.HandleFunc("PUT /{id}", bookmarkController.UpdateBookmark)
	
	return mux
}
