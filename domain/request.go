package domain

type SignupRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
