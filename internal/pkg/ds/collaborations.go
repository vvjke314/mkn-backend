package ds

import "github.com/google/uuid"

type Collaboration struct {
	Id        uuid.UUID `json:"id"`
	ProjectId uuid.UUID `json:"project_id"`
	UserId    uuid.UUID `json:"user_id"`
}
