package ds

import (
	"github.com/google/uuid"
)

type Project struct {
	id          uuid.UUID
	ownerId     uuid.UUID
	title       string
	description string
}
