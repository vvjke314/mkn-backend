package ds

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	id          uuid.UUID
	sectionId   uuid.UUID
	title       string
	description string
	deadline    time.Time
	status      string
	errorStatus int
}
