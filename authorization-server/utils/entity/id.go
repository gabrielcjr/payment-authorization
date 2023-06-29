package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func Parse(u string) (ID, error) {
	id, err := uuid.Parse(u)
	return id, err
}
