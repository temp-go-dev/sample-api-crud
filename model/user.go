package model

// User Model
type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus"`
	Version    int    `json:"version"`
}
