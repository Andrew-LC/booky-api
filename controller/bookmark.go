package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookmark-api/model"
	"bookmark-api/domain"
	"bookmark-api/service"
	middleware "bookmark-api/middlewares"
	util "bookmark-api/utils"
)

type BookmarkController struct {
	service service.BookmarkServiceInterface
}

func NewBookmarkController(s service.BookmarkServiceInterface) *BookmarkController {
	return &BookmarkController{service: s}
}


func (bc *BookmarkController) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var input domain.BookmarkRequest

	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	data, err := util.ExtractData(input.URL)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Failed to fetch URL metadata")
		return
	}

	bookmark := &model.Bookmark{
		UserID: userID,
		URL:    input.URL,
		Title:  data.Title,
		Notes:  input.Notes,
		Image:  data.Image,
		Tags:   input.Tags,
	}

	result, err := bc.service.CreateBookmark(r.Context(), bookmark)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to create bookmark")
		return
	}

	WriteJSON(w, http.StatusCreated, result)
}


func (bc *BookmarkController) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	bookmarks, err := bc.service.GetBookmarks(r.Context(), userID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to fetch bookmarks")
		return
	}

	WriteJSON(w, http.StatusOK, bookmarks)
}


func (bc *BookmarkController) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	var input domain.BookmarkupdateRequest

	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid bookmark ID")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid JSON body")
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

	updated, err := bc.service.UpdateBookmark(r.Context(), userID, uint(id), updates)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to update bookmark")
		return
	}

	WriteJSON(w, http.StatusOK, updated)
}


func (bc *BookmarkController) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r)
	if !ok {
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid bookmark ID") 
		return
	}

	if err := bc.service.DeleteBookmark(r.Context(), userID, uint(id)); err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to delete bookmark")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Bookmark deleted successfully",
	})
}

