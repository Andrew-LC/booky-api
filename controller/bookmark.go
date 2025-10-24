package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bookmark-api/model"
	middleware "bookmark-api/middlewares"
	util "bookmark-api/utils"
)

func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	type createBookmarkRequest struct {
		URL   string   `json:"url"`
		Notes string   `json:"notes"`
		Tags  []string `json:"tags"`
	}
	var input createBookmarkRequest
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	data, err := util.ExtractData(input.URL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch url", http.StatusBadRequest)
		return
	}

	bookmark := model.Bookmark{
		UserID: userID,
		URL:    input.URL,
		Title:  data.Title,
		Notes:  input.Notes,
		Image:   string(data.Image),
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
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bookmarks, err := model.GetBookmarks(userID)
	if err != nil {
		http.Error(w, "Failed to fetch bookmarks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmarks)
}


func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	type updateBookmarkRequest struct {
		URL   *string   `json:"url"`
		Title *string   `json:"title"`
		Notes *string   `json:"notes"`
		Tags  *[]string `json:"tags"`
	}
	var input updateBookmarkRequest
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Bookmark ID required", http.StatusBadRequest)
		return
	}
	var id uint
	fmt.Sscanf(idStr, "%d", &id)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{})
	if input.URL != nil {
		updates["url"] = *input.URL
	}
	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Notes != nil {
		updates["notes"] = *input.Notes
	}
	if input.Tags != nil {
		updates["tags"] = *input.Tags
	}

	updated, err := model.UpdateBookmark(userID, id, updates)
	if err != nil {
		http.Error(w, "Failed to update bookmark", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}


func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
