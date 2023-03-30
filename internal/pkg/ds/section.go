package ds

import "github.com/google/uuid"

type Section struct {
	id        uuid.UUID
	projectId uuid.UUID
	title     string
	color     string
}
