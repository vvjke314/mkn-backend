package ds

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID `json:"id"`
	OwnerId     uuid.UUID `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	LastEdited  time.Time `json:"last_edited"`
	Color       string    `json:"color"`
}
