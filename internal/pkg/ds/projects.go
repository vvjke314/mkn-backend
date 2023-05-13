package ds

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID `json:"id" gorm:"id"`
	OwnerId     uuid.UUID `json:"owner_id" gorm:"owner_id"`
	Title       string    `json:"title" gorm:"title" binding:"required"`
	Description string    `json:"description" gorm:"description" binding:"required"`
	LastEdited  time.Time `json:"last_edited" db:"last_edited" gorm:"type:timestamp"`
	Color       string    `json:"color" gorm:"color" binding:"required"`
}
