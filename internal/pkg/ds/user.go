package ds

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	username  string
	isManager bool
	email     string
}
