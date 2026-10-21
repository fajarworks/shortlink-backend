package model

type User struct {
	ID       int     `db:"id"`
	Email    string  `db:"email"`
	Password string  `db:"password"`
	FullName *string `db:"fullname"`
	ImageURL *string `db:"image_url"`
}
