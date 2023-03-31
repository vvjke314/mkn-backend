package ds

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id          uuid.UUID `json:"id"`
	SectionId   uuid.UUID `json:"section_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
	ErrorStatus int       `json:"error_status"`
}
