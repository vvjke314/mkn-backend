package ds

import "github.com/google/uuid"

type FavoriteProjects struct {
	id        uuid.UUID
	projectId uuid.UUID
	userId    uuid.UUID
}
