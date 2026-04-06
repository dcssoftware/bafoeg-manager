package models

import "github.com/google/uuid"

type IDModel struct {
	ID uuid.UUID `db:"id"`
}
