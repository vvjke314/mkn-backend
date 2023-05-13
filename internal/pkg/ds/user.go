package ds

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	IsManager int       `json:"is_manager" gorm:"is_manager"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"password" gorm:"password"`
}
