package ds

import "github.com/google/uuid"

type Collaborations struct {
	id        uuid.UUID
	projectId uuid.UUID
	userId    uuid.UUID
}
