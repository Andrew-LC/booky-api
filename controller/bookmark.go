package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bookmark-api/model"
)

func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var input struct {
		url     string    `json:"email"`
		title   string    `json:"title"`
		Notes   string    `json:"notes"`
		Tags    []string  `json:"string"`
	}
	userIDVal := r.Context().Value("userID")
	if userIDVal == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDVal.(uint)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	bookmark := model.Bookmark{
		UserID: userID,
		URL:    input.url,
		Title:  input.title,
		Notes:  input.Notes,
		Tags:   input.Tags,
	}

	if err := bookmark.CreateBookmark(); err != nil {
		http.Error(w, "Failed to create bookmark", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "bookmark successfully created"})
}

func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	if userIDVal == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDVal.(uint)

	bookmarks, err := model.GetBookmarks(userID)
	if err != nil {
		http.Error(w, "Failed to fetch bookmarks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmarks)
}


func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	if userIDVal == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDVal.(uint)

	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Bookmark ID required", http.StatusBadRequest)
		return
	}
	var id uint
	fmt.Sscanf(idStr, "%d", &id)

	if err := model.DeleteBookmark(userID, id); err != nil {
		http.Error(w, "Failed to delete bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bookmark %d deleted successfully", id)
}
