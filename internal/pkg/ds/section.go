package ds

import "github.com/google/uuid"

type Section struct {
	Id        uuid.UUID `json:"id"`
	ProjectId uuid.UUID `json:"project_id"`
	Title     string    `json:"title"`
	Color     string    `json:"color"`
}
