package main

import (
	"fmt"
	"net/http"
	"bookmark-api/middlewares"
	"bookmark-api/db"
	"bookmark-api/model"
	"bookmark-api/routes"
)



func main() {
	db.Connect()
	db.GetDB().AutoMigrate(&model.User{}, &model.Bookmark{})

	mux := http.NewServeMux()
	
	authMux := routes.AuthRoutes()
	bookmarkMux := routes.BookmarkRoutes()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/bookmark/", http.StripPrefix("/bookmark", middleware.JWTMiddleware(bookmarkMux)))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
	}

	fmt.Println("Server running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
