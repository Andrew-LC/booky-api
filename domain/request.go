package domain

import "github.com/lib/pq"

type SignupRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type BookmarkRequest struct {
	URL   string   `json:"url"`
	Notes string   `json:"notes"`
	Tags  []string `json:"tags"`
}

type BookmarkupdateRequest struct {
	URL   *string   `json:"url"`
	Title *string   `json:"title"`
	Notes *string   `json:"notes"`
	Tags  pq.StringArray `json:"tags"`
}
