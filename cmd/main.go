package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"bookmark-api/db"
	"bookmark-api/model"
	"bookmark-api/routes"
	"time"
)



func main() {
	db.Connect()
	db.GetDB().AutoMigrate(&model.User{}, &model.Bookmark{})

	mux := http.NewServeMux()
	routes.Register(mux, db.GetDB())
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		fmt.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server failed: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	}
	fmt.Println("Server gracefully stopped")
}
