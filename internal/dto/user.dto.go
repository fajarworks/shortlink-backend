package dto

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email    string  `json:"email"`
	FullName *string `json:"fullname"`
	ImageURL *string `json:"image_url"`
}
