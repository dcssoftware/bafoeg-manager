package models

import (
	"time"

	"github.com/google/uuid"
)

type EaktenFileModel struct {
	ID               string
	Source           uuid.UUID
	VorgangID        uuid.UUID
	SourceXdomeaFile bool
	FileID           uuid.UUID
	Created          time.Time
}
