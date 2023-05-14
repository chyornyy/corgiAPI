package models

type User struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	DateRegistrated string `json:"date_registrated"`
}
