package ds

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IsManager int       `json:"is_manager"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
